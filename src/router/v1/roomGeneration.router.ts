
import express from "express";
 
import { RoomgnerationJobSchema } from "../../dtos/roomgeneration.dto";
import { generateRoomsHandler } from "../../controller/roomgeneration";

import { validateRequestBody } from "../../validator";

const roomGenerationRouter=express.Router();
roomGenerationRouter.post('/',validateRequestBody(RoomgnerationJobSchema),generateRoomsHandler)

export default roomGenerationRouter