import { logger } from "../config/logger.config";
import Hotel from "../db/models/hotel";
import { NotFoundError } from "../utils/errors/app.error";
import BaseRepository from "./base.repositry";

// export async function createHotel(hotelData: createHotelDTO) {
//     const hotel = await Hotel.create(hotelData);

//     logger.info(`Hotel created: ${hotel.id}`);

//     return hotel;
// }

// export async function getHotelById(id: number){
//     const hotel = await Hotel.findByPk(id);

//     if(!hotel){
//         logger.error(`Hotel not found: ${id}`);
//         throw new NotFoundError("Hotel not found");
//     }

//     logger.info(`Hotel found: ${hotel.id}`);

//     return hotel;
// }

// export async function getAllHotels(){
//     const hotels = await Hotel.findAll({
//         where: {
//             deletedAt: null
//         }
//     });

//     if(!hotels){
//         logger.error("Hotels not found!");
//         throw new NotFoundError("No Hotels found!");
//     }

//     return hotels
// }

// export async function softDelete(id: number){
//     const deletedHotel = await Hotel.findByPk(id);

//     if(!deletedHotel){
//         logger.error(`No hotel found ${id}`);
//         throw new NotFoundError("No hotel found!");
//     }

//     deletedHotel.deletedAt = new Date;
//     await deletedHotel.save();

//     logger.info(`Hotel deleted successfully ${deletedHotel}`);

//     return deletedHotel;
// }


export class HotelRepository extends BaseRepository<Hotel>{
    constructor(){
        super(Hotel);
    }

    async findAll(){
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

    async softDelete(id: number){
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
}