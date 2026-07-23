import express  from "express";
import { bookingHandler } from "../../controllers/booking.controller";




const bookingRouter = express.Router();



bookingRouter.post("/", bookingHandler);


export default bookingRouter;