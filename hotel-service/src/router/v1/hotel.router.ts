
import express from "express";

import { createHotelHandler, getAllHotelsHandler, getHotelByIdHandler,deleteHotelHandler } from "../../controller/hotel.controller";
import { hotelSchema } from "../../validator/hotel.validator";
import { validateRequestBody } from "../../validator";

const hotelRouter=express.Router();
hotelRouter.post('/',
    validateRequestBody(hotelSchema),
    createHotelHandler);
hotelRouter.get('/:id',getHotelByIdHandler);
hotelRouter.get('/',getAllHotelsHandler);
hotelRouter.delete('/:id', deleteHotelHandler);

export default hotelRouter