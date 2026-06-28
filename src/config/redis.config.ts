import { serverConfig } from '.'

function connectToRedis() {
   try {
      return {
         port: serverConfig.REDIS_PORT,
         host: serverConfig.REDIS_HOST,
         maxRetriesPerRequest: null as null
      }
   } catch (error) {
      console.error('error connect to redis', error)
      throw error
   }
}

export const getRedisConnObject = connectToRedis()