import express  from "express";
import { validateQueryParam, validateRequestBody } from "../../validators";
import { roomSchema, updateRoomAvailabilitySchema } from "../../validators/room.validator";
import { getAvailableRoomsHandler, updateRoomAvailabilityHandler } from "../../controllers/room.controller";


const roomRouter = express.Router();

/**
 * @route POST /api/v1/rooms/getAvailableRooms
 * @desc Start the room availability extension scheduler
 * @access Public
 */

roomRouter.get("/getAvailableRooms", validateQueryParam(roomSchema), getAvailableRoomsHandler);
roomRouter.post("/update-rooms-id", validateRequestBody(updateRoomAvailabilitySchema), updateRoomAvailabilityHandler)

export default roomRouter;