export type getRoomAvailableDTO = {
    roomCategoryId: number,
    checkInDate: string,
    checkOutDate: string
}

export type updateRoomAvailabilityDTO = {
    bookingId: number,
    roomIds: number[]
}