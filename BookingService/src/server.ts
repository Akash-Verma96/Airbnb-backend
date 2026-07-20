import express from 'express'
import { serverConfig } from './config';
import v1Router from './routers/v1/index.router';
import v2Router from './routers/v2/index.router';
import { genericErrorHandler } from './middleware/error.middleware';
import { logger } from './config/logger.config';
import { attachCorrelationMiddleware } from './middleware/correlation.middleware';


const app = express();
app.use(express.json()); // Middleware to parse JSON request bodies used for serialization and deserialization of data in the request body
app.use(express.text());

app.use(attachCorrelationMiddleware); // middleware to attach correlation id to logger AsyncLocalStorage context

app.use("/api/v1", v1Router); // Registering all the routes from pingRouter to the app
app.use("/api/v2", v2Router); // Registering all the routes from pingRouter to the app


app.use(genericErrorHandler);

app.listen(serverConfig.PORT, () => {
    logger.info(`server is running on http://localhost:${serverConfig.PORT}`);
})
