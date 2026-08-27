import { IdempotencyKey, Prisma } from "../generated/prisma/client";
import { prismaClient } from "../lib/prisma"
import { BadRequestError, NotFoundError } from "../utils/errors/app.error";
import { validate as isValidUUID } from 'uuid';

/*
    createBooking  -- Done
    createIdempotencyKey  -- Done
    getIdempotencykey   -- Done
    getBookingById  -- Done
    confirmBooking  -- Done
    cancelBooking   -- Done
    finalizIdempotencyKey   -- Done

*/

export async function createBooking(bookingInput: Prisma.BookingCreateInput){

    const booking = await prismaClient.booking.create({
        data: bookingInput,
    })

    return booking;
}

export async function createIdempotencyKey(key: string, bookingId: number){
    const idempotencyKey = await prismaClient.idempotencyKey.create({
        data: {
            idemKey: key,
            booking: {
                connect: {
                    id: bookingId
                }
            }
        }
    });

    return idempotencyKey;
}

export async function getIdempotencyKeyWithLock(tx: Prisma.TransactionClient, key: string ){

    if(!isValidUUID(key)){
        throw new BadRequestError("Invalid idempotency key format");
    }

    const idempotencyKey: Array<IdempotencyKey> = await tx.$queryRaw
        `SELECT * FROM IdempotencyKey WHERE idemKey = ${key} FOR UPDATE`
    

    if(!idempotencyKey || idempotencyKey.length === 0){
        throw new NotFoundError("Idempotency key not fount");
    }

    return idempotencyKey[0];
}

export async function getBookingById(bookingId: number){
    const booking = await prismaClient.booking.findUnique({
        where: {
            id: bookingId
        }
    })

    return booking;
}

export async function confirmBooking(tx: Prisma.TransactionClient, bookingId: number){
    const booking = await tx.booking.update({
        where: {
            id: bookingId
        },
        data:{
            status: "CONFIRMED"
        }

    })

    return booking;
}

export async function cancelBooking(bookingId: number){
    const booking = await prismaClient.booking.update({
        where:{
            id: bookingId
        },
        data:{
            status: "CANCELLED"
        }
    })

    return booking;
}

export async function finalizeIdempotencyKey(tx: Prisma.TransactionClient, key: string) {
    const idempotencyKey = await tx.idempotencyKey.update({
        where:{
            idemKey: key
        },
        data:{
            finalized: true
        }
    })

    return idempotencyKey;
}



