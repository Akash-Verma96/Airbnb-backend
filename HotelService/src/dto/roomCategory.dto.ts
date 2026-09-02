import { RoomType } from "../db/models/roomCategory"

export type createRoomCategroyDTO = {
    hotelId: number,
    roomType: RoomType,
    roomNo: number,
    price: number
}