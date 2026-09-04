import z from 'zod'

export const roomSchema = z.object({
    roomCategoryId: z.string({message: "Room Category Id must be present"}),
    checkInDate: z.string({message: "CheckIn Date must be present"}),
    checkOutDate: z.string({message: "Checkout Date must be present"})
})

export const updateRoomAvailabilitySchema = z.object({
    bookingId: z.number({message: "Booking id must be present"}),
    roomIds: z.array(z.number({message: "Room ids must be present"})),
})