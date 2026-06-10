import { PrismaMariaDb } from "@prisma/adapter-mariadb";
import { PrismaClient } from "@prisma/client";

const adapter = new PrismaMariaDb(process.env.DATABASE_URL as string);

export default new PrismaClient({ adapter });
