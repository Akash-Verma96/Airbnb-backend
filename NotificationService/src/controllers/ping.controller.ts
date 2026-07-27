import { NextFunction, Request, Response } from 'express';
// import fs from 'fs/promises';
import { NotFoundError} from '../utils/errors/app.error';
import { logger } from '../config/logger.config';


export const pingHandler = async (req: Request, res: Response, next: NextFunction) => {
    try {
        // await fs.readFile("sample");

        logger.info("inside pingHandler");
        res.status(200).json({
            message : 'pong'
        });
    } catch (error) {
        // throw new InternalServerError("Interval Server Error");
        throw new NotFoundError("File Not Found!");
    }
}