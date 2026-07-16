import { NextFunction, Request, Response } from 'express';
// import fs from 'fs/promises';
import { NotFoundError} from '../utils/errors/app.error';
import { logger } from '../config/logger.config';
import User from '../db/models/user.models';


export const pingHandler = async (req: Request, res: Response, next: NextFunction) => {
    try {
        // await fs.readFile("sample");

        logger.info("inside pingHandler");

        // const akash = await User.create({ firstName: 'Akash', lastName: 'Verma' });
        // console.log("Jane's auto-generated ID:", akash.id);

        const users = await User.findAll();


        res.status(200).json({
            user: users,
            // _id: akash.id,
            message : 'pong'
        });
    } catch (error) {
        // throw new InternalServerError("Interval Server Error");
        throw new NotFoundError("File Not Found!");
    }
}