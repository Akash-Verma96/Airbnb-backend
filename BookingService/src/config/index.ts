// this file contains all the basic configuration logic for the app server to work
import dotenv from 'dotenv';

type ServerConfig = {
    PORT: number,
    REDIS_SERVER_URL: string,
    LOCK_TTL: number,
    HOTEL_SERVICE_URL: string
}

function loadEnv(){
    dotenv.config();
    console.log("Env loaded successfully");
}

loadEnv();

export const serverConfig: ServerConfig = {
    PORT: Number(process.env.PORT) || 3001, // process.env.PORT is a string, so we need to convert it to a number
    REDIS_SERVER_URL: process.env.REDIS_SERVER_URL || "redis://localhost:6379",
    LOCK_TTL: Number(process.env.LOCK_TTL) || 5000,
    HOTEL_SERVICE_URL: process.env.HOTEL_SERVICE_URL || "http://localhost:3002"
}