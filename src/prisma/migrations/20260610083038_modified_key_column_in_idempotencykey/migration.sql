/*
  Warnings:

  - You are about to drop the column `key` on the `IdempotencyKey` table. All the data in the column will be lost.

*/
-- DropIndex
DROP INDEX `IdempotencyKey_key_key` ON `IdempotencyKey`;

-- AlterTable
ALTER TABLE `IdempotencyKey` DROP COLUMN `key`,
    ADD COLUMN `idemKey` VARCHAR(191) NOT NULL DEFAULT '';
