export type createBookingDto = {
    userId: number,
    hotelId: number,
    bookingAmount: number,
    totalGuests: number,
    checkInDate: string,
    checkOutDate: string,
    roomCategoryId: number
}