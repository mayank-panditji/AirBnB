import { pingHandler } from "../../controller/ping.controller";
import express from "express";
import {validateRequestBody} from "../../validator";
import { pingschema } from "../../validator/ping.validator";
const pingRouter=express.Router();

pingRouter.get('/',validateRequestBody(pingschema),pingHandler);
pingRouter.get('/health',(req,res)=>{
    res.status(200).send('ok');
})

export default pingRouter