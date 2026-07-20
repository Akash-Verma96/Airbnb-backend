import { NextFunction, Request, Response } from "express";
import {v4 as uuidV4} from "uuid";
import { asyncLocalStorage } from "../utils/helpers/request.helpers";


export const attachCorrelationMiddleware = (req : Request,res: Response, next: NextFunction) => {

    const correlationId = uuidV4();

    asyncLocalStorage.run({ correlationid : correlationId},() => {
        next();
    })
    // asyncLocalStorage.enterWith({ correlationid : correlationId});
    // next();
    
}