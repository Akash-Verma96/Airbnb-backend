import winston from "winston"
// import { MongoDB } from "winston-mongodb";
import { getCorrelationId } from "../utils/helpers/request.helpers";
import DailyRotateFile from "winston-daily-rotate-file";

export const logger = winston.createLogger({
    format: winston.format.combine(
        winston.format.timestamp({ format: "MM-DD-YY HH-MM-SS" }),
        winston.format.json(),
        winston.format.printf(({level, message, timestamp, ...data}) => {
            const output = {
                level,
                message,
                timestamp,
                correlationId : getCorrelationId(),
                data};

            return JSON.stringify(output);
        })
    ),
    transports: [
        new winston.transports.Console(),
        new DailyRotateFile({
            filename: "logs/%DATE%-app.log",
            datePattern: "YYYY-MM-DD",
            maxSize: "20m",
            maxFiles: "14d"
        })
    ]
})




// const ConnectDB = async ()=>{
//     logger.info('Connecting to database...');
//     const MongoClient =  require('mongodb').MongoClient;
//     const url = process.env.MONGO_URI;

//     const client = new MongoClient(url);
//     await client.connect();





//     const transportOptions = {
//         db: await Promise.resolve(client),
//         collection: 'logger'
//     };

//     logger.add(new MongoDB(transportOptions));
 

//     logger.info('Connected to database.',{url});
// }

// ConnectDB();


// following entry should appear in logger collection and will contain
// metadata JSON-property containing url field
