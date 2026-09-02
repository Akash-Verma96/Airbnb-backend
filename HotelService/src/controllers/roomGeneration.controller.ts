import { StatusCodes } from "http-status-codes";
import { generateRooms } from "../services/roomGeneration.service";
import { Request, Response } from "express";

export async function generateRoomHandler(req: Request, res: Response){
    const generatedRooms = await generateRooms(req.body);

    res.status(StatusCodes.OK).json({
        message:"Hotel Room updated Successfully",
        data: generatedRooms,
        success: true
    })
}