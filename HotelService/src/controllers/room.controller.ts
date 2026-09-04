import { Request, Response } from "express"
import { getAvailableRoomsService, updateRoomsAvailabilityService } from "../services/room.service";
import { StatusCodes } from "http-status-codes";



export async function getAvailableRoomsHandler(req: Request, res: Response){
   
       const availableRooms = await getAvailableRoomsService({
        roomCategoryId:  Number(req.query.roomCategoryId),
        checkInDate: req.query.checkInDate as string,
        checkOutDate: req.query.checkOutDate as string
       });
   
       res.status(StatusCodes.CREATED).json({
           message: "Rooms Fetched Successfully!",
           data: availableRooms,
           success: true,
       })
   }

export async function updateRoomAvailabilityHandler(req: Request, res: Response){
    const updatedRooms = await updateRoomsAvailabilityService(req.body);
   
       res.status(StatusCodes.CREATED).json({
           message: "Rooms updated Successfully!",
           data: updatedRooms,
           success: true,
       })

}
