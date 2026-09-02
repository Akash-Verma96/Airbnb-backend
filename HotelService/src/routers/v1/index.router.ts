import express from 'express';
import pingRouter from '../v1/ping.router';
import hotelRouter from './hotel.router';
import roomGenerationRuter from './roomGeneration.router';

const v1Router = express.Router();

v1Router.use('/', pingRouter);
v1Router.use('/hotels', hotelRouter);
v1Router.use('/', roomGenerationRuter);

export default v1Router;