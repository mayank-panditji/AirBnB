import Room from "../db/models/room";
import BaseRepository from "./base.repository";
import { CreationAttributes } from "sequelize";

export class RoomRepository extends BaseRepository<Room> {
    constructor() {
        super(Room);
    }

    async findByRoomCategoryIdAndDate(roomCategoryId: number, currentDate: Date) {
        return await this.model.findOne({
            where: {
                roomCategoryId,
                dateOfAvailability: currentDate,
                deletedAt: null
            }
        });
    }

    async getNextRoomNo(hotelId: number): Promise<number> {
        const maxRoom = await this.model.findOne({
            where: { hotelId },
            order: [["roomNo", "DESC"]],
        });
        return maxRoom ? (maxRoom as any).roomNo + 1 : 1;
    }

    async bulkCreate(rooms: CreationAttributes<Room>[]) {
        return await this.model.bulkCreate(rooms);
    }
}