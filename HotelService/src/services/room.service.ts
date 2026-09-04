import { getRoomAvailableDTO, updateRoomAvailabilityDTO } from "../dto/room.dto";
import { RoomRepository } from "../repositories/room.repository";


const roomRepository = new RoomRepository();

export async function getAvailableRoomsService(roomsData : getRoomAvailableDTO){
    return await roomRepository.findbyRoomCategoryIdAndDateRange(roomsData.roomCategoryId,roomsData.checkInDate,roomsData.checkOutDate);
}

export async function updateRoomsAvailabilityService(updateRoomAvailabilityDTO: updateRoomAvailabilityDTO){
    return await roomRepository.updateRoomsAvailability(updateRoomAvailabilityDTO.bookingId, updateRoomAvailabilityDTO.roomIds);
}