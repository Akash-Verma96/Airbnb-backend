import { RoomType } from "../db/models/roomCategory"

export type createRoomCategroyDTO = {
    hotel_id: number,
    roomType: RoomType,
    roomNo: number,
    price: number
}