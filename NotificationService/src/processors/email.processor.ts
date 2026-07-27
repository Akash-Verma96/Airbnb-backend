import { Job, Worker } from "bullmq";
import { NotificationDto } from "../dto/notification.dto";
import { MAILER_QUEUE } from "../queues/mailer.queue";
import { getRedisConnObject } from "../config/redis.config";
import { MAILER_PAYLOAD } from "../producers/email.producer";
import { sendEmail } from "../services/mailer.service";
import { renderMailTemplate } from "../handlebars/template.handlebar";
import { logger } from "../config/logger.config";


export function setupMailerWorker(){
    const emailWorker = new Worker<NotificationDto>(
        MAILER_QUEUE, // queue name
        async(job: Job) =>{

            if(job.name != MAILER_PAYLOAD){
                throw new Error("Invalid job name");
            }

            console.log("Processing the mail job ....!");

            const payload = job.data;

            console.log("Job payload: ", JSON.stringify(payload));

            // call the service layer from here

            const emailContent = await renderMailTemplate(payload.templateId, payload.params);


            await sendEmail(payload.to, payload.subject, emailContent);

            logger.info(`Email sent to ${payload.to} with subject "${payload.subject}"`);
        },
        {
            connection: getRedisConnObject()
        }
    )

    emailWorker.on("failed", () => {
        console.log("Email processing failed!");
    })

    emailWorker.on("completed", ()=>{
        console.log("Email processing completed successfully!");
    })
}