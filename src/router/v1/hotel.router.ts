
import express from "express";

import { createHotelHandler, getAllHotelsHandler, getHotelByIdHandler } from "../../controller/hotel.controller";
import { hotelSchema } from "../../validator/hotel.validator";
import { validateRequestBody } from "../../validator";

const hotelRouter=express.Router();
hotelRouter.post('/',
    validateRequestBody(hotelSchema),
    createHotelHandler);
hotelRouter.get('/:id',getHotelByIdHandler);
hotelRouter.get('/',getAllHotelsHandler);

export default hotelRouter