import { Job, Worker } from "bullmq";
import { RoomGenerationJob } from "../dto/roomGeneration.dto";
import { ROOM_GENERATION_QUEUE } from "../queues/roomGeneration.queue";
import { getRedisConnObject } from "../config/redis.config";
import { ROOM_GENERATION_PAYLOAD } from "../producers/roomGeneration.producer";
import { generateRooms } from "../services/roomGeneration.service";
import { logger } from "../config/logger.config";


export function setupRoomGenerationWorker(){
    const roomGenerationWorker = new Worker<RoomGenerationJob>(
        ROOM_GENERATION_QUEUE, // queue name
        async(job: Job) =>{

            if(job.name != ROOM_GENERATION_PAYLOAD){
                throw new Error("Invalid job name");
            }

            const payload = job.data;
            console.log(`Room generation started processing job ${JSON.stringify(payload)}`);

            const roomsData = await generateRooms(payload);

            logger.info(`Total rooms created ${roomsData.totalRoomsCreated} and total date processed ${roomsData.totalDatesProcessed}`);
           
        },
        {
            connection: getRedisConnObject()
        }
    )

    roomGenerationWorker.on("failed", () => {
        console.log("Room Generation processing failed!");
    })

    roomGenerationWorker.on("completed", ()=>{
        console.log("Room Generation processing completed successfully!");
    })
}