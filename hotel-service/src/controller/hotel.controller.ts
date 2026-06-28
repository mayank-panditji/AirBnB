import { Request, Response, NextFunction } from 'express';
import { createHotelService, getAllHotelsService, getHotelByIdService, deleteHotelService } from '../services/hotel.service';
import { StatusCodes } from 'http-status-codes';

export async function createHotelHandler(req: Request, res: Response ,next: NextFunction){
 const hotelResponse = await createHotelService(req.body);
 res.status(StatusCodes.CREATED).json({
    message: "Hotel created successfully",
    data: hotelResponse,
    success: true,
 })
}
export async function getHotelByIdHandler(req: Request, res: Response ,next: NextFunction){
    const hotelResponse = await getHotelByIdService(Number(req.params.id));
    res.status(StatusCodes.OK).json({
       message: "Hotel retrieved successfully",
       data: hotelResponse,
       success: true,
    })
   }
   export async function getAllHotelsHandler(req: Request, res: Response ,next: NextFunction){
      const hotelsResponse = await getAllHotelsService();
      res.status(StatusCodes.OK).json({
         message: "Hotels retrieved successfully",
         data: hotelsResponse,
         success: true,
      })

   }
   export async function updateHotelHandler(req: Request, res: Response ,next: NextFunction){  
    res.status(StatusCodes.NOT_IMPLEMENTED);
    }
    export async function deleteHotelHandler(req: Request, res: Response ,next: NextFunction){
        const hotelsResponse = await deleteHotelService(Number(req.params.id));
      res.status(StatusCodes.OK).json({
         message: "Hotels soft deleted successfully",
         data: hotelsResponse,
         success: true,
      })
    }
