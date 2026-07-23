import { Prisma } from "../generated/prisma/client";
import { prisma } from "../lib/prisma"

export async function createBooking(bookingInput: Prisma.BookingCreateInput){
    const booking = await prisma.booking.create({
        data: bookingInput,
    })

    return booking;
}