import { Job, Worker } from "bullmq";
import { NotificationDto } from "../dto/notification.dto";
import { MAILER_QUEUE } from "../queues/mailer.queue";
import { getRedisConnObject } from "../config/redis.config";
import { MAILER_PAYLOAD } from "../producers/email.producer";
import nodemailer from 'nodemailer'

function sendMail(payload: NotificationDto){
    const transporter = nodemailer.createTransport({
        host: 'smtp.gmail.com',
        port: 465,
        secure: true, // use SSL
        auth: {
            user: process.env.SMTP_USER,
            pass: process.env.SMTP_PASSWORD,
        }
    })

    const mailOptions = {
        from: 'akt64725@gmail.com',
        to: payload.to,
        subject: payload.subject,
        text: payload.params.name
    };

    // Send the email
    transporter.sendMail(mailOptions, function(error, info){
    if (error) {
        console.log('Error:', error);
    } else {
        console.log('Email sent:', info.response);
    }
    });
}


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

            sendMail(payload);


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