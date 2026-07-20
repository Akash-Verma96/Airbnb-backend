import express  from "express";
import { createHotelHandler, softDeleteHandler, getAllHotelHandler, getHotelByIdHandler, updateHotelNameByIdHandler } from "../../controllers/hotel.controller";
import { validateRequestBody } from "../../validators";
import { hotelSchema } from "../../validators/hotel.validator";


const hotelRouter = express.Router();


hotelRouter.post(
    '/',
    validateRequestBody(hotelSchema),
    createHotelHandler);

    
hotelRouter.get('/getAllHotels', getAllHotelHandler);

hotelRouter.get('/:id', getHotelByIdHandler);

hotelRouter.delete('/:id', softDeleteHandler);

hotelRouter.patch('/', updateHotelNameByIdHandler);


export default hotelRouter;