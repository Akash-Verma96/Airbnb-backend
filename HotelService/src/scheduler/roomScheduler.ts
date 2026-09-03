import * as cron from 'node-cron'
import { logger } from '../config/logger.config';
import { serverConfig } from '../config';
import { RoomRepository } from '../repositories/room.repository';
import { RoomCategoryRepository } from '../repositories/roomCategory.repository';
import { RoomGenerationJob } from '../dto/roomGeneration.dto';
import { addRoomGenerationToQueue } from '../producers/roomGeneration.producer';

const roomRepository = new RoomRepository();
const roomCategoryRepository = new RoomCategoryRepository();

let cronJob: cron.ScheduledTask | null = null;

export const startScheduler = (): void => {
    if(cronJob) {
        logger.warn("Room scheduler is already running");
        return;
    }

    cronJob = cron.schedule(serverConfig.ROOM_CRON, async () => {
        try {
            logger.info("Starting room availability extension check!");
            await extendRoomAvailability();
            logger.info("Room availability extension check completed!");
        } catch (error) {
            logger.info("Error in room availability extension scheduler:", error);   
        }
    }, {
        timezone: 'UTC'
    });

    cronJob.start();
    logger.info(`Room availability extension scheduler started - running every ${serverConfig.ROOM_CRON}`);
};

// stop room availability extension
export const stopScheduler = (): void => {
    if(cronJob) {
        cronJob.stop();
        cronJob = null;
        logger.info("Room availability extension scheduler stopped!");
    }
};

// Get scheduler status
export const getSchedulerStatus = (): {isRunning: boolean} => {
    return {
        isRunning: cronJob !== null && cronJob.getStatus() === 'scheduled'
    };
}

 const extendCategoryAvailability = async (roomCategoryId: number, latestDate: Date): Promise<void> => {
        try {
            const nextDate = new Date(latestDate);
            nextDate.setDate(nextDate.getDate() + 1);


            // check if the room category still exists

            const roomCategory = await roomCategoryRepository.findById(roomCategoryId);

            if(!roomCategory){
                logger.warn(`Room category ${roomCategoryId} not found, skipping extension`);
                return;
            }


            const existingRoom = await roomRepository.findByRoomCategoryIdAndDate(roomCategoryId, nextDate);
            if(existingRoom) {
                logger.debug(`Room for category ${roomCategoryId} on ${nextDate.toISOString()} already exists, skipping`);
                return;
            }

            const endDate = new Date(nextDate);

            endDate.setDate(endDate.getDate() + 1);

            const jobDate: RoomGenerationJob = {
                roomCategoryId: roomCategoryId,
                startDate: nextDate.toISOString(),
                endDate: endDate.toISOString(),
                priceOverride: roomCategory.price,
                batchSize: 1
            };

            // now we have all payload -> add job to queue
            await addRoomGenerationToQueue(jobDate)

            logger.info(`Added room generation job for category ${roomCategoryId} on ${nextDate.toISOString()}`);

        } catch (error) {
            logger.error(`Error extending availability for room category ${roomCategoryId}:`, error);
        }
    }

// extend Room availability for 1 day for all categories so that booking made 
const extendRoomAvailability = async (): Promise<void> => {
    try {
        const roomCategoriesWithLatestDates = await roomRepository.findLatestDatesForAllCategories();


        if (roomCategoriesWithLatestDates.length === 0) {
            logger.info('No room categories found with availability dates');
            return;
        }

        logger.info(`Found ${roomCategoriesWithLatestDates.length} room categories to extend`);

        // Now we have all categories which i have to process to make it available for booking
        // Processing each room category -- lage raho kar sakte ho 

        for(const categoryData of roomCategoriesWithLatestDates) {
            await extendCategoryAvailability(categoryData.roomCategoryId,categoryData.latestDate);
        }

    } catch (error) {
        logger.error('Error extending room availability:', error);
        throw error;
    }

    // now extend availability for a specific room category -- sab category ka availabal karana hoga

   
}


export const manualExtendAvailability = async (): Promise<void> => {
    logger.info('Manual room availability extension triggered');
    await extendRoomAvailability();
}; 