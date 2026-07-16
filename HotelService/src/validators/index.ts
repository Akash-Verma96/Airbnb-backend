import { NextFunction, Request, Response } from "express";
import {  z } from "zod";
import { logger } from "../config/logger.config";

export const validateRequestBody = (schema: z.ZodObject) => {
    return async (req: Request, res: Response, next: NextFunction) =>{
        try {
            logger.info("Request Body is Validating");
            await schema.parseAsync(req.body);
            logger.info("Request Body is Validated");
            next();
        } catch (error) {
            return res.status(400).json({ 
                message: "Invalid request body",
                success: false,
                error: error
            });
        }
    }
}

export const validateQueryParam = (schema: z.ZodObject) => {
    return async (req: Request, res: Response, next: NextFunction) =>{
        try {
            await schema.parseAsync(req.query);
            console.log("Request Query is Valid :", req.query);
            next();
        } catch (error) {
            return res.status(400).json({ 
                message: "Invalid request query",
                success: false,
                error: error
            });
        }
    }
}