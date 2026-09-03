import express  from "express";
import { validateRequestBody } from "../../validators";
import { generateRoomHandler } from "../../controllers/roomGeneration.controller";
import { RoomGenerationJobSchema } from "../../dto/roomGeneration.dto";


const roomGenerationRuter = express.Router();

/**
 * @route POST /api/v1/generateRooms
 * @desc Generates the rooms for first time mannually
 * @access Public
 */
roomGenerationRuter.post('/generateRooms', validateRequestBody(RoomGenerationJobSchema),  generateRoomHandler)


export default roomGenerationRuter;