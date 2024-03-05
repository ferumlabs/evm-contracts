import dotenv from "dotenv";
export const DEV = "development";
export const PROD = "production";
export const NODE = "node";
export const ETHERNAL_NODE = "ethernal-node";
export const env = process.env.ENV || DEV;
dotenv.config({
  path: `.env.${env}`,
});
