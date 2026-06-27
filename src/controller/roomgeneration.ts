import { Request, Response, NextFunction } from "express";
import {generateRooms} from "../services/roomgeneration.servics"
import { RoomGenerationJob } from "../dtos/roomgeneration.dto";
import logger from "../config/logger";

export async function generateRoomsController(
  req: Request,
  res: Response,
  next: NextFunction
) {
  try {
    const { roomCategoryId, startDate, endDate, batchSize, priceOverride } = req.body;

    const jobData: RoomGenerationJob = {
      roomCategoryId,
      startDate,
      endDate,
      batchSize,
      priceOverride,
    };

    logger.info(`Received room generation request for category ${roomCategoryId}`);

    const result = await generateRooms(jobData);

    return res.status(200).json({
      success: true,
      message: "Rooms generated successfully",
      data: result,
    });
  } catch (error) {
    next(error); // tumhare global error-handling middleware ko bhej diya, jo NotFoundError/BadRequestError handle karta hoga
  }
}