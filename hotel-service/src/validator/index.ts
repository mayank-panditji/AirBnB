import { ZodObject } from "zod";
import { Request,Response,NextFunction } from "express";
import logger from "../config/logger";
export const validateRequestBody=(schema:ZodObject)=>{
    return async(req:Request,res:Response,next:NextFunction)=>{
        try {
            logger.info('Validating request body');
           await schema.parseAsync(req.body);
           logger.info('Request body is valid');
            next();
        } catch (error) {
            logger.error('Request body is invalid');
            res.status(400).json({
                message:'Invalid request body',
               success:false,
                error:error
            });

        }
    }
}
export const validateQueryParams=(schema:ZodObject)=>{
return async(req:Request,res:Response,next:NextFunction)=>{
        try {
           await schema.parseAsync(req.query);
           logger.info('Query parameters are valid');
            next();
        } catch (error) {
            logger.error('Query parameters are invalid');
            res.status(400).json({
                message:'Invalid request body',
               success:false,
                error:error
            });

        }
    }
};
