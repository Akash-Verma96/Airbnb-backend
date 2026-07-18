import express  from "express";
import { pingHandler } from "../../controllers/ping.controller";
import { validateRequestBody } from "../../validators";
import { pingSchema } from "../../validators/ping.validator";



const pingRouter = express.Router();


// pingRouter.get("/ping",validateRequestBody(pingSchema), validateQueryParam(querySchema), pingHandler);
pingRouter.get("/ping",validateRequestBody(pingSchema), pingHandler);
// pingRouter.post("/ping",validateRequestBody(pingSchema), pingHandler);


export default pingRouter;