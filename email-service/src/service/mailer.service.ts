import  logger  from "../config/logger";
import { serverConfig } from "../config";
import transporter from "../config/mailer.config";
import { InternalServerError } from "../utils/errors/app.error";

export async function sendemail(to:string,subject:string,body:string) {
   try{
     await transporter.sendMail({
        from:serverConfig.MAIL_USER,
        to,
        subject,
        html:body
    })
    logger.info(`email sent to ${to} with subject:${subject}`)
   }catch(error){
    throw new InternalServerError(`faild to send email`)
   }
}