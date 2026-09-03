import { Router } from "express";
import {
    startSchedulerHandler,
    stopSchedulerHandler,
    getSchedulerStatusHandler,
    manualExtendAvailabilityHandler
} from "../../controllers/roomScheduler.controller";

const roomSchederRouter = Router();

/**
 * @route POST /api/v1/scheduler/start
 * @desc Start the room availability extension scheduler
 * @access Public
 */
roomSchederRouter.post("/start", startSchedulerHandler);

/**
 * @route POST /api/v1/scheduler/stop
 * @desc Stop the room availability extension scheduler
 * @access Public
 */
roomSchederRouter.post("/stop", stopSchedulerHandler);

/**
 * @route GET /api/v1/scheduler/status
 * @desc Get the current status of the room availability extension scheduler
 * @access Public
 */
roomSchederRouter.get("/status", getSchedulerStatusHandler);

/**
 * @route POST /api/v1/scheduler/extend
 * @desc Manually trigger room availability extension
 * @access Public
 */
roomSchederRouter.post("/extend", manualExtendAvailabilityHandler);

export default roomSchederRouter; 