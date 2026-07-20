import { Request, Response, NextFunction } from "express"
import { createHotelService, softDeleteService, getAllHotelService, getHotelByIdService, updateHotelNameByIdService } from "../services/hotel.service"
import { StatusCodes } from "http-status-codes";


export async function createHotelHandler(req: Request, res: Response, next: NextFunction){
    // call service layer

    const hotelResponse = await createHotelService(req.body);

    // send the response

    res.status(StatusCodes.CREATED).json({
        message: "Hotel created Successfully!",
        data: hotelResponse,
        success: true,
    })
}

export async function getHotelByIdHandler(req: Request, res: Response, next: NextFunction) {

    const hotelResponse = await getHotelByIdService(Number(req.params.id));

    res.status(StatusCodes.OK).json({
        message: "Hotel found Successfully!",
        data: hotelResponse,
        success: true,
    })
}

export async function getAllHotelHandler(req: Request, res: Response, next: NextFunction) {
    const hotelsResponse = await getAllHotelService();


    res.status(StatusCodes.ACCEPTED).json({
        message: "Hotel Detail Found!",
        data: hotelsResponse,
        success: true
    })
}

export async function softDeleteHandler(req: Request,res: Response, next: NextFunction){
    const deletedHotel = await softDeleteService(Number(req.params.id));

    res.status(StatusCodes.OK).json({
        message: "Hotel Deleted Successfully",
        data: deletedHotel,
        success: true
    })
}

export async function updateHotelNameByIdHandler(req: Request, res: Response, next: NextFunction){
    const updateHotel = await updateHotelNameByIdService(req.body);

    res.status(StatusCodes.ACCEPTED).json({
        message: "Hotel Updated Successfully",
        data: updateHotel,
        success: true
    })
}