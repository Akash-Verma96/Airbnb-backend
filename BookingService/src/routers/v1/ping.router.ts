import express  from "express";
import { pingHandler } from "../../controllers/ping.controller";
import { validateRequestBody } from "../../validators";
import { pingSchema } from "../../validators/ping.validators.schema";



const pingRouter = express.Router();


// pingRouter.get("/ping",validateRequestBody(pingSchema), validateQueryParam(querySchema), pingHandler);
pingRouter.get("/ping",validateRequestBody(pingSchema), pingHandler);


export default pingRouter;