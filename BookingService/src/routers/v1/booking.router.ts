import express  from "express";
import { confirmBookingHandler, createBookingHandler, getAllBookingsHandler } from "../../controllers/booking.controller";
import { validateRequestBody } from "../../validators";
import { createBookingSchema } from "../../validators/booking.validators";




const bookingRouter = express.Router();



bookingRouter.post('/',validateRequestBody(createBookingSchema), createBookingHandler);
bookingRouter.post('/confirm/:idempotencyKey', confirmBookingHandler);


/**
 * @route POST /api/v1/booking/getAllBookings
 * @desc Gets all your bookings
 * @access Public
 */

bookingRouter.get('/getAllBookings/:id', getAllBookingsHandler)


export default bookingRouter;