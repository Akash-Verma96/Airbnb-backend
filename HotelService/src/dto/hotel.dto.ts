import { RoomType } from "../db/models/roomCategory";

export type createHotelDTO = {
    name: string;
    address: string;
    location: string;
    price: number;
    roomType: RoomType;
    hostId: number;
    rating?: number;
    ratingCount?: number;
}

export type updateHotelDTO = {
    name?: string;
    address?: string;
    location?: string;
    price?: number;
    room_type?: RoomType;
}