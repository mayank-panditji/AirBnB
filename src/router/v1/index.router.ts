import  express from "express";
import pingRouter from "./ping.router";
import hotelRouter from "./hotel.router";
import roomGenerationRoutes from "./roomGeneration.router";
const v1Router=express.Router();
v1Router.use('/ping',pingRouter)
v1Router.use('/hotel',hotelRouter)
v1Router.use("/room-generation", roomGenerationRoutes);
v1Router.use("/room-generation", roomGenerationRoutes);
export default v1Router