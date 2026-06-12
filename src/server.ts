import express from "express";
import {serverConfig} from "./config";
import pingRouter from "./router/v1/ping.router";
import v1Router from "./router/v1/index.router";
import v2Router from "./router/v2/index.router";
import z from "zod";
import logger from "./config/logger";
import { genericErrorHandler } from "./middleware/error.middleware";
import { attachCorrelationIdMiddleware } from "./middleware/correlation.middleware";
import { setUpMailerWorker } from "./processor/email.processor";
import { NotificationDto } from "./dto/notification.dto";
import { addEmailtoQueue } from "./producer/email.producer";
const app = express();
app.use(express.json())


app.use(attachCorrelationIdMiddleware)


app.use('/api/v1',v1Router)
app.use('/api/v2',v2Router)

app.use(genericErrorHandler)

app.listen(serverConfig.PORT , () => {
  logger.info(`server is running at ${serverConfig.PORT}`);
  console.log("hello");
  setUpMailerWorker();
  logger.info(`mailer worker setup completed`)
  logger.info("Server started successfully",{data:"some additional data"});
  const smapleNotification:NotificationDto={
    to:"sample",
    subject:"Sample email",
    templateid:"sample-tempplate",
    params:{
      name:"Mayank Trivedi",
      orederId:"12345"
    }
  }
  addEmailtoQueue(smapleNotification)
});