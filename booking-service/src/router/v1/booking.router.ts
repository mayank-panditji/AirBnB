
import express from "express";
import {validateRequestBody} from "../../validator";
import { createBookingschema } from "../../validator/booking.validator";
import { confirmBookingHandler, createBookingHandler } from "../../controller/booking.controller";
const bookingRouter=express.Router();

bookingRouter.post('/',validateRequestBody(createBookingschema),createBookingHandler);
bookingRouter.post('/confirm/:idempotencyKey', confirmBookingHandler);
export default bookingRouter