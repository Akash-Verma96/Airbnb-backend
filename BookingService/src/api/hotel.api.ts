import axios from "axios"
import { serverConfig } from "../config"

export async function getAvailableRooms(roomCategoryId: number, checkInDate: string, checkOutDate: string){
    const response = await axios.get(`${serverConfig.HOTEL_SERVICE_URL}/api/v1/rooms/getAvailableRooms`, {
        params: {
            roomCategoryId,
            checkInDate,
            checkOutDate
        },
    });

    return response.data;
}


export async function updateRoomAvailability(bookingId: number, roomIds: number[]){
    const response = await axios.post(`${serverConfig.HOTEL_SERVICE_URL}/api/v1/rooms/update-rooms-id`, {
            bookingId,
            roomIds
    })

    return response.data;
}