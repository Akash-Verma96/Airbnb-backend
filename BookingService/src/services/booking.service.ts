import { Prisma } from "../generated/prisma/client";
import { createBooking } from "../repositories/booking.repository";

export async function createBookingService(bookingInput: Prisma.BookingCreateInput){
    const booking = await createBooking(bookingInput);

    return booking;
}

export async function finlizeBookingService(){
    // TODO
}