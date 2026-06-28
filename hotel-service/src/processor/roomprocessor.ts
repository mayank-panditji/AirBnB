import { Job, Worker } from "bullmq";
import { ROOM_GENERATION_QUEUE } from "../queue/roomgeneration";
import { RoomGenerationJob } from "../dtos/roomgeneration.dto";
import { getRedisConnObject } from "../config/redis.config";
import {ROOM_GENERATION_PAYLOAD} from "../producer/roomgeneration_producer"
import logger  from "../config/logger";
import { generateRooms } from "../services/roomgeneration.servics";
export const setUpRoomGenerationWorker=()=>{
    const roomGenerationProcessor=new Worker<RoomGenerationJob>(
    ROOM_GENERATION_QUEUE,async(job:Job)=>{
            if(job.name!=ROOM_GENERATION_PAYLOAD){
                throw new Error("Invalid job name")
            }
            const payload=job.data;
            console.log(`Processing room generation for:${JSON.stringify(payload)}`);
            await generateRooms(payload);
              logger.info(`room generation completed for:${JSON.stringify(payload)}`);
    },
    {
        connection:getRedisConnObject
    }
)
roomGenerationProcessor.on("failed",()=>{
    console.error("room generation  processing failed")
})
roomGenerationProcessor.on("completed",()=>{
    console.error("room generation  processing succescfully")
})
}