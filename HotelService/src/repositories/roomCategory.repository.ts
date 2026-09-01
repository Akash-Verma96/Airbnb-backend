import BaseRepository from "./base.repositry";
import RoomCategory from "../db/models/roomCategory";


export class RoomCategoryRepository extends BaseRepository<RoomCategory>{
    constructor(){
        super(RoomCategory);
    }

    async findRoomCategoryByHotelId(hotelId: number){
        const roomCategory = await this.model.findAll({
            where:{
                hotelId: hotelId,
                deletedAt: null
            }
        })
            

            return roomCategory;
        }
}