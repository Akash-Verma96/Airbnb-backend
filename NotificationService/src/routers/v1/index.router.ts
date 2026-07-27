import express from 'express';
import pingRouter from '../v2/ping.router';

const v1Router = express.Router();

v1Router.use("/", pingRouter);

export default v1Router;