import { Request, Response, NextFunction } from "express"
import { createHotelService, getHotelByIdService } from "../services/hotel.service"

export async function createHotelHandler(req: Request, res: Response, next: NextFunction){
    // call service layer

    const hotelResponse = await createHotelService(req.body);

    // send the response

    res.status(200).json({
        message: "Hotel created Successfully",
        data: hotelResponse,
        success: true,
    })
}

export async function getHotelByIdHandler(req: Request, res: Response, next: NextFunction) {
    const hotelResponse = await getHotelByIdService(Number(req.params.id));

    res.status(200).json({
        message: "Hotel found",
        data: hotelResponse,
        success: true,
    })
}