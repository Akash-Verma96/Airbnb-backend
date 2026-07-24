/*
  Warnings:

  - Made the column `bookingId` on table `idempotencykey` required. This step will fail if there are existing NULL values in that column.

*/
-- DropForeignKey
ALTER TABLE `idempotencykey` DROP FOREIGN KEY `IdempotencyKey_bookingId_fkey`;

-- AlterTable
ALTER TABLE `idempotencykey` MODIFY `bookingId` INTEGER NOT NULL;

-- AddForeignKey
ALTER TABLE `IdempotencyKey` ADD CONSTRAINT `IdempotencyKey_bookingId_fkey` FOREIGN KEY (`bookingId`) REFERENCES `Booking`(`id`) ON DELETE RESTRICT ON UPDATE CASCADE;
