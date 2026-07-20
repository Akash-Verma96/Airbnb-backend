import { createHotelDTO, updateHotelNameDTO } from "../dto/hotel.dto";
import { createHotel, softDelete, getAllHotels, getHotelById, updateHotelNameById } from "../repositories/hotel.repository";

export async function createHotelService(hotelData: createHotelDTO){
    const hotel = await createHotel(hotelData);

    return hotel;
}

export async function getHotelByIdService(id: number) {
    const hotel = await getHotelById(id);

    return hotel;
}

export async function getAllHotelService(){
    const hotels = await getAllHotels();

    return hotels;
}

export async function softDeleteService(id : number){
    const deletedHotel = await softDelete(id);

    return deletedHotel;
}

export async function updateHotelNameByIdService(updateData: updateHotelNameDTO){
    const updatedhotel = await updateHotelNameById(updateData);

    return updatedhotel;
}