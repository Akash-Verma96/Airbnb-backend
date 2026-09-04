import { CreationAttributes, Op } from "sequelize";
import Room from "../db/models/room"
import BaseRepository from "./base.repositry"


export class RoomRepository extends BaseRepository<Room>{
    constructor(){
        super(Room);
    }

    async findByRoomCategoryIdAndDate(
        roomCategoryId: number,
        currentDate: Date
    ) {
        return await this.model.findOne({
            where: {
                roomCategoryId,
                dateOfAvailability: currentDate,
                deletedAt: null
            }
        })
    }

    async bulkCreate(rooms: CreationAttributes<Room>[]){
        return this.model.bulkCreate(rooms)
    }

    async findLatestDateByRoomCategoryId(roomCategoryId: number): Promise<Date | null> {
        const result = await this.model.findOne({
            where: {
                roomCategoryId,
                deletedAt: null
            },
            attributes: ['dateOfAvailability'],
            order: [['dateOfAvailability', 'DESC']]
        })

        return result ? result.dateOfAvailability : null;
    }


    async findLatestDatesForAllCategories(): Promise<Array<{roomCategoryId: number, latestDate: Date}>> {
        const results = await this.model.findAll({
            where: {
                deletedAt: null
            },

            attributes: [
                'roomCategoryId',
                [this.model.sequelize!.fn('MAX', this.model.sequelize!.col('date_of_availability')), 'latestDate']
            ],
            group: ['roomCategoryId'],
            raw: true
        });

        return results.map((result: any) => ({
            roomCategoryId: result.roomCategoryId,
            latestDate: new Date(result.latestDate)
        }));
    }

    async findbyRoomCategoryIdAndDateRange(
        roomCategoryId: number,
        checkInDate: string,
        checkOutDate: string
    ){
        return await this.model.findAll({
            where: {
                roomCategoryId,
                bookingId: null,
                dateOfAvailability: {
                    [Op.between]: [checkInDate, checkOutDate]
                }
            }
        })
    }

    async updateRoomsAvailability(bookingId: number, roomIds: number[]){
        return await this.model.update(
            {bookingId},
            { where: {id: { [Op.in]: roomIds }}}
        )
    }
}