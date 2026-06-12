import { Job, Worker } from "bullmq";
import { NotificationDto } from "../dto/notification.dto";
import { MAILER_QUEUE } from "../queues/mailer.queue";
import { getRedisConnObject } from "../config/redis.config";
import { MAILER_PAYLOAD } from "../producer/email.producer";
export const setUpMailerWorker=()=>{
    const emailProcessor=new Worker<NotificationDto>(
    MAILER_QUEUE,async(job:Job)=>{
            if(job.name!=MAILER_PAYLOAD){
                    throw new Error("Invalid job name")
            }
            const payload=job.data;
            console.log(`processing email for:${JSON.stringify(payload)}`);
            
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