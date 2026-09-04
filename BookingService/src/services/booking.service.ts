import { getAvailableRooms, updateRoomAvailability } from "../api/hotel.api";
import { serverConfig } from "../config";
import { redlock } from "../config/redis.config";
import { createBookingDto } from "../dto/booking.dto";
import { prismaClient } from "../lib/prisma";
import { confirmBooking, createBooking, createIdempotencyKey, finalizeIdempotencyKey, getIdempotencyKeyWithLock } from "../repositories/booking.repository";
import { BadRequestError, InternalServerError, NotFoundError } from "../utils/errors/app.error";
import { generateIdempotencyKey } from "../utils/generateIdempotencyKey";

type AvailableRoom = {
    id: number,
    roomCategoryId: number,
    dateOfAvailability: Date
}


export async function createBookingService(createBookingDto: createBookingDto){

    const bookingResourse = `hotel:${createBookingDto.hotelId}`;
    const ttl = serverConfig.LOCK_TTL;

    try {
        await redlock.acquire([bookingResourse],ttl);
        // modifies the room resource id
        const availableRooms = await getAvailableRooms(createBookingDto.roomCategoryId,createBookingDto.checkInDate,createBookingDto.checkOutDate);
        console.log(availableRooms)
        const checkOutDate = new Date(createBookingDto.checkOutDate);
        const checkInDate = new Date(createBookingDto.checkInDate);

        const totlaNights = Math.ceil((checkOutDate.getTime() - checkInDate.getTime()) / (1000 * 3600 * 24));


        if( availableRooms.data.length == 0 || totlaNights > availableRooms.data.length){
            throw new BadRequestError('No rooms available for given dates');
        }


        const booking = await createBooking({
            userId: createBookingDto.userId,
            hotelId: createBookingDto.hotelId,
            totalGuests: createBookingDto.totalGuests,
            bookingAmount: createBookingDto.bookingAmount,
            roomCategoryId: createBookingDto.roomCategoryId,
            checkInDate: new Date(createBookingDto.checkInDate),
            checkOutDate: new Date(createBookingDto.checkOutDate)
        });

        const idempotencyKey = generateIdempotencyKey();

        await createIdempotencyKey(idempotencyKey,booking.id);
        
        await updateRoomAvailability(booking.id, availableRooms.data.map((room: AvailableRoom) => room.id));
    
        return {
            bookingId: booking.id,
            idempotencyKey: idempotencyKey
        };
        
    } catch (error: any) {
        if (error.name === 'BadRequestError' || error.statusCode === 400) {
            throw error;
        }
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
        // update the booking Id in room if booking gets fail

        return booking;
    })
}