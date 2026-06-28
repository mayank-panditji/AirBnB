import { Job, Worker } from "bullmq";
import { NotificationDto } from "../dto/notification.dto";
import { MAILER_QUEUE } from "../queues/mailer.queue";
import { getRedisConnObject } from "../config/redis.config";
import { MAILER_PAYLOAD } from "../producer/email.producer";
import { renderMailTemplate } from "../templates/templates.handler";
import { sendemail } from "../service/mailer.service";
import logger  from "../config/logger";
export const setUpMailerWorker=()=>{
    const emailProcessor=new Worker<NotificationDto>(
    MAILER_QUEUE,async(job:Job)=>{
            if(job.name!=MAILER_PAYLOAD){
                    throw new Error("Invalid job name")
            }
            const payload=job.data;
            console.log(`processing email for:${JSON.stringify(payload)}`);
             const emailContent=await renderMailTemplate(payload.templateid,payload.params)
             await sendemail(payload.to,payload.subject,emailContent)
             logger.info(`email sent to ${payload.to} with subject ${payload.subject}`)
    },
    {
        connection:getRedisConnObject
    }
)
emailProcessor.on("failed",()=>{
    console.error("email processing failed")
})
emailProcessor.on("completed",()=>{
    console.error("email processing succescfully")
})
}