// import "dotenv/config";
// import { PrismaMariaDb } from "@prisma/adapter-mariadb";
// import { PrismaClient } from "../generated/prisma/client";

// const adapter = new PrismaMariaDb({
//   host: process.env.DATABASE_HOST || 'localhost',
//   user: process.env.DATABASE_USER  || 'root',
//   password: process.env.DATABASE_PASSWORD || '1234',
//   database: process.env.DATABASE_NAME || 'test',
//   connectionLimit: 5,
// });
// const prismaClient = new PrismaClient({ adapter });
// export { prismaClient };

import "dotenv/config";
import { PrismaMariaDb } from "@prisma/adapter-mariadb";
import { PrismaClient } from "../generated/prisma/client";

const adapter = new PrismaMariaDb({
  host: process.env.DATABASE_HOST || '',
  user: process.env.DATABASE_USER || '',
  password: process.env.DATABASE_PASSWORD || '',
  database: process.env.DATABASE_NAME || '',
  port: Number(process.env.DATABASE_PORT) || 10337, 
  connectionLimit: 5,
  
  ssl: {
    rejectUnauthorized: false
  },
});

const prismaClient = new PrismaClient({ adapter });
export { prismaClient };