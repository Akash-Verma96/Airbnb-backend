import { StatusCodes } from "http-status-codes";
// import { generateRooms } from "../services/roomGeneration.service";
import { Request, Response } from "express";
import { addRoomGenerationToQueue } from "../producers/roomGeneration.producer";

export async function generateRoomHandler(req: Request, res: Response){
    await addRoomGenerationToQueue(req.body)

    res.status(StatusCodes.OK).json({
        message:"Hotel Room updated Successfully",
        data: {},
        success: true
    })
}