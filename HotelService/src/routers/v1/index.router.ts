import express from 'express';
import pingRouter from '../v1/ping.router';
import hotelRouter from './hotel.router';
import roomGenerationRuter from './roomGeneration.router';
import roomSchederRouter from './roomScheduler.router';
import roomRouter from './room.router';

const v1Router = express.Router();

v1Router.use('/', pingRouter);
v1Router.use('/hotels', hotelRouter);
v1Router.use('/', roomGenerationRuter);
v1Router.use('/scheduler', roomSchederRouter);
v1Router.use('/rooms', roomRouter);

export default v1Router;