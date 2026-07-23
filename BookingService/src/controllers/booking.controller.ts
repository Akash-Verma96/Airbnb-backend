import { NextFunction, Request, Response } from 'express';
import { InternalServerError} from '../utils/errors/app.error';
import { createBookingService } from '../services/booking.service';
import { StatusCodes } from 'http-status-codes';
// import { createBookingService } from '../services/booking.service';
// import { prisma } from "../lib/prisma";



export const bookingHandler = async (req: Request, res: Response, next: NextFunction) => {
    try {
        const booking = await createBookingService(req.body);

        return res.status(StatusCodes.OK).json({
            "message" : "Booking created Successfully!",
            "data" : booking
        })
    } catch (error) {
        console.log(error);
        throw new InternalServerError("Internal Server ERROR");
    }
}