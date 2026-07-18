import express from 'express';
import pingRouter from './ping.router';

const v2Router = express.Router();

v2Router.use("/", pingRouter);

export default v2Router;