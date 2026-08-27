import { serverConfig } from "../config";
import { redlock } from "../config/redis.config";
import { createBookingDto } from "../dto/booking.dto";
import { prismaClient } from "../lib/prisma";
import { confirmBooking, createBooking, createIdempotencyKey, finalizeIdempotencyKey, getIdempotencyKeyWithLock } from "../repositories/booking.repository";
import { BadRequestError, InternalServerError, NotFoundError } from "../utils/errors/app.error";
import { generateIdempotencyKey } from "../utils/generateIdempotencyKey";



export async function createBookingService(createBookingDto: createBookingDto){

    const bookingResourse = `hotel:${createBookingDto.hotelId}`;
    const ttl = serverConfig.LOCK_TTL;

    try {
        await redlock.acquire([bookingResourse],ttl);


        const booking = await createBooking({
            userId: createBookingDto.userId,
            hotelId: createBookingDto.hotelId,
            totalGuests: createBookingDto.totalGuests,
            bookingAmount: createBookingDto.bookingAmount
        });

        const idempotencyKey = generateIdempotencyKey();

        await createIdempotencyKey(idempotencyKey,booking.id);

        return {
            bookingId: booking.id,
            idempotencyKey: idempotencyKey
        };
        
    } catch (error) {
        console.error("Error acquiring lock: ", error);
        throw new InternalServerError("Error already lock acquired by other person.");
    }
}

// explore potential issues in this service race condition-- Fixed
export async function confirmBookingService(idempotencyKey: string){
    return await prismaClient.$transaction(async (tx) =>{
        const idempotencyKeyData = await getIdempotencyKeyWithLock(tx,idempotencyKey);

        if(!idempotencyKeyData){
            throw new NotFoundError("Idempotency key not found");
        }

        if(idempotencyKeyData.finalized){
            throw new BadRequestError("Idempotency key already finalized")
        }

        const booking = await confirmBooking(tx,idempotencyKeyData.bookingId);

        await finalizeIdempotencyKey(tx,idempotencyKey);

        return booking;
    })
}