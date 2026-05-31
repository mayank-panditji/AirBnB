import express from "express";
import {serverConfig} from "./config";
import pingRouter from "./router/v1/ping.router";
import v1Router from "./router/v1/index.router";
import v2Router from "./router/v2/index.router";
import z from "zod";
import logger from "./config/logger";
import { genericErrorHandler } from "./middleware/error.middleware";
import { attachCorrelationIdMiddleware } from "./middleware/correlation.middleware";
const app = express();
app.use(express.json())


app.use(attachCorrelationIdMiddleware)


app.use('/api/v1',v1Router)
app.use('/api/v2',v2Router)

app.use(genericErrorHandler)

app.listen(serverConfig.PORT , () => {
  logger.info(`server is running at ${serverConfig.PORT}`);
  console.log("hello");
  logger.info("Server started successfully",{data:"some additional data"});
});