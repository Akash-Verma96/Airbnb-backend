import { mailerQueue } from "../queues/mailer.queue";
import { NotificationDto } from "../dto/notification.dto";


export const MAILER_PAYLOAD = "payload:mail";

export const addEmailToQueue = async (payload: NotificationDto) =>{
    await mailerQueue.add(MAILER_PAYLOAD,payload);
    console.log(`Email added to queue: ${JSON.stringify(payload)}`);
}


