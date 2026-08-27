import  { Redis }  from 'ioredis';
import { serverConfig } from '.';
import Redlock from 'redlock';


// singleton pattern to create redis client
function connectToRedis(){
    try {
        let connection: Redis;

       
        return () => {
            if(!connection){
                connection = new Redis(serverConfig.REDIS_SERVER_URL);
;
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


export const redlock = new Redlock([getRedisConnObject()],{
    
    driftFactor: 0.01, // multiplied by lock ttl to determine drift time

    // The max number of times Redlock will attempt to lock a resource
    // before erroring.
    retryCount: 10,

    // the time in ms between attempts
    retryDelay: 200, // time in ms

    // the max time in ms randomly added to retries
    // to improve performance under high contention
    // see https://www.awsarchitectureblog.com/2015/03/backoff.html
    retryJitter: 200,
})