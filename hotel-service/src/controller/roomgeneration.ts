import { Request, Response, NextFunction } from "express";
import {generateRooms} from "../services/roomgeneration.servics"
import { RoomGenerationJob } from "../dtos/roomgeneration.dto";
import logger from "../config/logger";
import { addRoomGenerationJobtoQueue } from "../producer/roomgeneration_producer";
import { StatusCodes } from "http-status-codes";

export async function generateRoomsHandler(
  req: Request,
  res: Response,
) {
  await addRoomGenerationJobtoQueue(req.body)

    

     res.status(StatusCodes.OK).json({
      success: true,
      message: "Rooms generated job added to queue",
      data: {},
    });  
  }
