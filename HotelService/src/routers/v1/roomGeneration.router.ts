import express  from "express";
import { validateRequestBody } from "../../validators";
import { generateRoomHandler } from "../../controllers/roomGeneration.controller";
import { RoomGenerationJobSchema } from "../../dto/roomGeneration.dto";


const roomGenerationRuter = express.Router();


roomGenerationRuter.post('/generateRooms', validateRequestBody(RoomGenerationJobSchema),  generateRoomHandler)


export default roomGenerationRuter;