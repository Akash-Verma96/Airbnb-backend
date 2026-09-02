import { roomGenerationQueue } from "../queues/roomGeneration.queue";
import { RoomGenerationJob } from "../dto/roomGeneration.dto";


export const ROOM_GENERATION_PAYLOAD = "payload:roomGeneration";

export const addRoomGenerationToQueue = async (payload: RoomGenerationJob) =>{
    await roomGenerationQueue.add(ROOM_GENERATION_PAYLOAD,payload);
    console.log(`Room generation added to queue: ${JSON.stringify(payload)}`);
}


