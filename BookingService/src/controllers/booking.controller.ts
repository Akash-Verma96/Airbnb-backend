import { Request, Response } from 'express';
import { confirmBookingService, createBookingService } from '../services/booking.service';
import { StatusCodes } from 'http-status-codes';


interface userParams {
    idempotencyKey: string;
}

export const createBookingHandler = async (req: Request, res: Response) => {

    const booking = await createBookingService(req.body);

    return res.status(StatusCodes.OK).json({
        bookingId: booking.bookingId,
        IdempotencyKey: booking.idempotencyKey
    })
}

export const confirmBookingHandler = async (req: Request<userParams>, res: Response) => {
    const booking = await confirmBookingService(req.params.idempotencyKey);

    res.status(StatusCodes.OK).json({
        bookingId: booking.id,
        status: booking.status,
    });
}
