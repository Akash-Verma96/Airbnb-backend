import express  from "express";
import { createHotelHandler, softDeleteHandler, getAllHotelHandler, getHotelByIdHandler, updateHotelByIdHandler } from "../../controllers/hotel.controller";
import { validateRequestBody } from "../../validators";
import { hotelSchema } from "../../validators/hotel.validator";


const hotelRouter = express.Router();


hotelRouter.post('/',validateRequestBody(hotelSchema),createHotelHandler);

    
hotelRouter.get('/getAllHotels', getAllHotelHandler);

hotelRouter.get('/:id', getHotelByIdHandler);

hotelRouter.delete('/:id', softDeleteHandler);

hotelRouter.patch('/:id', updateHotelByIdHandler);

export default hotelRouter;