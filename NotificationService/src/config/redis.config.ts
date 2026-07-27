import Redis from "ioredis";
import { serverConfig } from ".";

// singleton pattern to create redis client
function connectToRedis(){
    try {
        let connection: Redis;

        const redisConfig = {
            host: serverConfig.REDIS_HOST,
            port: serverConfig.REDIS_PORT,
            maxRetriesPerRequest: null, // Disable automatic reconnection
        }

        return () => {
            if(!connection){
                connection = new Redis(redisConfig);
                return connection;
            }

            return connection;
        }

    } catch (error) {
        console.log("Error while connection redis", error);
        throw error;
    }
}

export const getRedisConnObject = connectToRedis();