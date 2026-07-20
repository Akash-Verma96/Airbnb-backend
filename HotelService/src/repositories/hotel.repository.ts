import { logger } from "../config/logger.config";
import Hotel from "../db/models/hotel";
import { createHotelDTO, updateHotelNameDTO } from "../dto/hotel.dto";
import { NotFoundError } from "../utils/errors/app.error";

export async function createHotel(hotelData: createHotelDTO) {
    const hotel = await Hotel.create(hotelData);

    logger.info(`Hotel created: ${hotel.id}`);

    return hotel;
}

export async function getHotelById(id: number){
    const hotel = await Hotel.findByPk(id);

    if(!hotel){
        logger.error(`Hotel not found: ${id}`);
        throw new NotFoundError("Hotel not found");
    }

    logger.info(`Hotel found: ${hotel.id}`);

    return hotel;
}

export async function getAllHotels(){
    const hotels = await Hotel.findAll({
        where: {
            deletedAt: null
        }
    });

    if(!hotels){
        logger.error("Hotels not found!");
        throw new NotFoundError("No Hotels found!");
    }

    return hotels
}

export async function softDelete(id: number){
    const deletedHotel = await Hotel.findByPk(id);

    if(!deletedHotel){
        logger.error(`No hotel found ${id}`);
        throw new NotFoundError("No hotel found!");
    }

    deletedHotel.deletedAt = new Date;
    await deletedHotel.save();

    logger.info(`Hotel deleted successfully ${deletedHotel}`);

    return deletedHotel;
}

export async function updateHotelNameById(updateData: updateHotelNameDTO) {
    const hotel = await Hotel.findByPk(updateData.id);

    if(!hotel){
        logger.error(`No hotel found ${updateData.id}`);
        throw new NotFoundError("Hotel Not found");
    }

    hotel.name = updateData.name;
    hotel.save();

    logger.info("Hotel Name updated successfully");

    return hotel;
}