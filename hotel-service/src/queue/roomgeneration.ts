import { Queue } from 'bullmq'
import { getRedisConnObject } from '../config/redis.config'

export const ROOM_GENERATION_QUEUE = "queue-room-generation"
export const roomGenrationQueue = new Queue(ROOM_GENERATION_QUEUE, {
    connection: getRedisConnObject,
})