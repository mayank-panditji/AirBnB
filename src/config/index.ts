import dotenv from 'dotenv'
type ServerConfig={
 PORT:number,
 REDIS_PORT:number,
 REDIS_HOST:string,
}
 function loadEnv()
{
    dotenv.config();
}

loadEnv();
export const serverConfig:ServerConfig={
    PORT:Number(process.env.PORT)|| 3001,
    REDIS_PORT:process.env.REDIS_PORT?Number(process.env.REDIS_PORT):6379,
    REDIS_HOST:process.env.REDIS_HOST||'localhost',
}