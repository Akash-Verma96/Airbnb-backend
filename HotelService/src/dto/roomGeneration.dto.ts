import { z } from 'zod'

/* TODO: Extend the controller to take request schema and decide whether it wants a sync or async flow

    - Create a controller and take payload from user and check wheather he wants immediate or scheduled from payload you can check
    - Based on the requirment call the service layer
    - call service startScheduler() for scheduled type 
    - call addRoomGenerationToQueue(req.body) for immediate generation
    - you can decide this inside roomGeneration/roomScheduler controller
*/
export const RoomGenerationRequestSchema = z.object({
    roomCategoryId: z.number().positive(),
    startDate: z.iso.datetime(),
    endDate: z.iso.datetime(),
    scheduleType: z.enum(['immediate', 'scheduled']).default('immediate'),
    scheduledAt: z.iso.datetime().optional(),
    priceOverride: z.number().positive().optional(),
});



export const RoomGenerationJobSchema = z.object({
    roomCategoryId: z.number().positive(),
    startDate: z.iso.datetime(),
    endDate: z.iso.datetime(),
    priceOverride: z.number().positive().optional(),
    batchSize: z.number().positive().default(100),
});


export type RoomGenerationJob = z.infer<typeof RoomGenerationJobSchema>;
export type RoomGenerationRequest = z.infer<typeof RoomGenerationRequestSchema>;

export interface RoomGenerationResponse {
    success: boolean;
    totalRoomsCreated: number;
    totalDatesProcessed: number;
    errors: string[];
    jobId: string;
}