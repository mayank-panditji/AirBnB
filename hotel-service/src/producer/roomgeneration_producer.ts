import { RoomGenerationJob } from "../dtos/roomgeneration.dto";
import { roomGenrationQueue } from "../queue/roomgeneration";
export const ROOM_GENERATION_PAYLOAD="payload:room-generation"
export const addRoomGenerationJobtoQueue=async(payload:RoomGenerationJob)=>{
    await roomGenrationQueue.add(ROOM_GENERATION_PAYLOAD,payload)
    
}