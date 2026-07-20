export type createHotelDTO = {
    name: string;
    address: string;
    location: string;
    rating?: number;
    ratingCount?: number;
}

export type updateHotelNameDTO = {
    id: number
    name: string
}