import { start } from "node:repl";
import {z} from "zod";

export const roomgenerationSchema=z.object({
   roomCategoryId:z.number().positive(),
   startDate:z.string().datetime(),
   endDate:z.string().datetime(),
   scheduleType:z.enum(['immediate','scheduled']).default('immediate'),
   scheduledAt:z.string().datetime().optional(),
   priceOverride:z.number().positive().optional(),
})
export const RoomgnerationJobSchema=z.object({
    roomCategoryId:z.number().positive(),
   startDate:z.string().datetime(),
   endDate:z.string().datetime(),
   priceOverride:z.number().positive().optional(),
   batchSize:z.number().positive().default(100),
})
export type RoomGenerationRequest=z.infer<typeof roomgenerationSchema>
export type RoomGenerationJob=z.infer<typeof RoomgnerationJobSchema>
export interface RoomGenerationResponse{
    success:boolean;
    totalRoomsCreated:number;
    totalDatesProcessed:number;
   errors:string[];
   jobId:string[];
}