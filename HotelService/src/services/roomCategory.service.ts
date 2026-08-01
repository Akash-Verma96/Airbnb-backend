import { createRoomCategroyDTO } from "../dto/roomCategory.dto";
import { HotelRepository } from "../repositories/hotel.repository";
import { RoomCategoryRepository } from "../repositories/roomCategory.repository";
import { NotFoundError } from "../utils/errors/app.error";


const roomCategoryRepository = new RoomCategoryRepository();
const hotelRepository = new HotelRepository();

export async function createRoomCategroyService(roomCategoryData: createRoomCategroyDTO){
    const roomDetail = await roomCategoryRepository.create(roomCategoryData);

    return roomDetail;
}

export async function getRoomCategoryByHotelIdService(hotelId: number) {
    // find hotel with hotel id
    const hotel = await hotelRepository.findById(hotelId);

    if(!hotel){
        throw new NotFoundError("Hotel Not found");
    }
    // find room category by hotelId and return

    const roomCategoryDetail = await roomCategoryRepository.findRoomCategoryByHotelId(hotelId);

    return roomCategoryDetail;
}

export async function deleteRoomCategoryService(id: number) {
    const isAvailableroomCategory = await roomCategoryRepository.findById(id);

    if(!isAvailableroomCategory){
        throw new NotFoundError(`Room Category with id ${id} not found!`);
    }

    await roomCategoryRepository.delete({id});

    return true;
}

