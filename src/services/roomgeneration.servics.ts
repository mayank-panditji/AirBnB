import { RoomCategoryRepository } from "../repositary/roomCategory.repository";
import { RoomGenerationJob } from "../dtos/roomgeneration.dto";
import { NotFoundError, BadRequestError } from "../utils/errors/app.error";
import RoomCategory from "../db/models/roomCategory";
import { RoomRepository } from "../repositary/room.repository";
import Room from "../db/models/room";
import logger from "../config/logger";
import { CreationAttributes } from "sequelize";
const roomCategoryRepository=new RoomCategoryRepository();
const roomRepository = new RoomRepository();
export async function generateRooms(jobData: RoomGenerationJob) {

  let totalRoomCreated=0;
  let totalDateProcessed=0;

    const roomCategory=await roomCategoryRepository.findById(jobData.roomCategoryId);
    if (!roomCategory) {
        throw new NotFoundError(`Room category with id ${jobData.roomCategoryId} not found`);
    }
   const startDate=new Date(jobData.startDate);
   const endDate=new Date(jobData.endDate);
  if(startDate>=endDate){
    throw new BadRequestError(`Start date is greater than end date`);
  }
  if(startDate<new Date()){
    throw new BadRequestError(`Start date is in the past`);
  }

  const totaldays=Math.ceil((endDate.getTime()-startDate.getTime())/(1000*60*60*24));
  
logger.info(`Generating rooms for room category  for ${totaldays} days`);

  const batchSize=jobData.batchSize || 100;
 const currDate=new Date(startDate);
 while(currDate<endDate){
   const batchEndDate=new Date(currDate);
   batchEndDate.setDate(batchEndDate.getDate()+batchSize);
   if(batchEndDate>endDate){
    batchEndDate.setTime(endDate.getTime());
   }
    const batchResult=await processDateBatch(roomCategory,currDate,batchEndDate,jobData.priceOverride);
    totalRoomCreated+=batchResult.roomsCreated;
    totalDateProcessed+=batchResult.datesProcessed;
    currDate.setTime(currDate.getTime()+batchSize);
 }
 return {
  totalRoomCreated,
  totalDateProcessed
 }



}
// use a better query to get a better roooms
//here n+1 problem generate
export async function processDateBatch(roomCategory:RoomCategory,startDate:Date,endDate:Date,priceOverride?:number){
  let roomsCreated=0;
  let datesProcessed=0;
  const roomToCreate:CreationAttributes<Room>[]=[];
  const currentDate=new Date(startDate);
  while(currentDate<=endDate){
    const existingRoom=await roomRepository.findByRoomCategoryIdAndDate(roomCategory.id,currentDate);
    if(!existingRoom){
       roomToCreate.push({
        hotelId:roomCategory.hotelId,
        roomCategoryId:roomCategory.id,
        dateOfAvailability:currentDate,
        price:priceOverride || roomCategory.price,
        createdAt:new Date(),
        updatedAt:new Date(),
        deletedAt:null
      });
      
    }
    currentDate.setDate(currentDate.getDate()+1);
    datesProcessed++;
  }
  if(roomToCreate.length>0){
    await roomRepository.bulkCreate(roomToCreate);
    roomsCreated+=roomToCreate.length;
  }
  return {roomsCreated,datesProcessed};
}