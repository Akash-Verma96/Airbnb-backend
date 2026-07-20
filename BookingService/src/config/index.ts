// this file contains all the basic configuration logic for the app server to work
import dotenv from 'dotenv';

type ServerConfig = {
    PORT: number
}

function loadEnv(){
    dotenv.config();
    console.log("Env loaded successfully");
}

loadEnv();

export const serverConfig: ServerConfig = {
    PORT: Number(process.env.PORT) || 3001 // process.env.PORT is a string, so we need to convert it to a number
}
