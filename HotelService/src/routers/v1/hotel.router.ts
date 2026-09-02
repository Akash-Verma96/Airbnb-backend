import express  from "express";
import { createHotelHandler, softDeleteHandler, getAllHotelHandler, getHotelByIdHandler, updateHotelNameByIdHandler } from "../../controllers/hotel.controller";
import { validateRequestBody } from "../../validators";
import { hotelSchema } from "../../validators/hotel.validator";
import { generateRoomHandler } from "../../controllers/roomGeneration.controller";


const hotelRouter = express.Router();


hotelRouter.post('/',validateRequestBody(hotelSchema),createHotelHandler);

    
hotelRouter.get('/getAllHotels', getAllHotelHandler);

hotelRouter.get('/:id', getHotelByIdHandler);

hotelRouter.delete('/:id', softDeleteHandler);

hotelRouter.patch('/', updateHotelNameByIdHandler);

hotelRouter.post('/generateRooms', generateRoomHandler)


export default hotelRouter;