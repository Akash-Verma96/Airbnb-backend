/*
  Warnings:

  - You are about to drop the column `idempotencyKeyId` on the `booking` table. All the data in the column will be lost.
  - A unique constraint covering the columns `[bookingId]` on the table `IdempotencyKey` will be added. If there are existing duplicate values, this will fail.

*/
-- DropForeignKey
ALTER TABLE `booking` DROP FOREIGN KEY `Booking_idempotencyKeyId_fkey`;

-- DropIndex
DROP INDEX `Booking_idempotencyKeyId_key` ON `booking`;

-- AlterTable
ALTER TABLE `booking` DROP COLUMN `idempotencyKeyId`;

-- AlterTable
ALTER TABLE `idempotencykey` ADD COLUMN `bookingId` INTEGER NULL,
    ADD COLUMN `finalized` BOOLEAN NOT NULL DEFAULT false;

-- CreateIndex
CREATE UNIQUE INDEX `IdempotencyKey_bookingId_key` ON `IdempotencyKey`(`bookingId`);

-- AddForeignKey
ALTER TABLE `IdempotencyKey` ADD CONSTRAINT `IdempotencyKey_bookingId_fkey` FOREIGN KEY (`bookingId`) REFERENCES `Booking`(`id`) ON DELETE SET NULL ON UPDATE CASCADE;
