
import { NextFunction, Request,Response} from "express"
import fs from 'fs/promises'
import { AppError, InternalServerError, NotFoundError } from "../utils/errors/app.error"
import logger from "../config/logger"
export const pingHandler=async(req:Request,res:Response,next:NextFunction)=>{
  logger.info('Received ping request');
    res.status(200).json({
        message: 'pong',
        
    })
   }
  