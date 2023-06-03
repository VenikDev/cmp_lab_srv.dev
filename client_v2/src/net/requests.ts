import ky from "ky";
import {HOST_V1} from "./consts";
import {Logger} from "../common/logger";
import {IError} from "../models/error";

export async function getAnalysis<T, TErr = IError>(
  key: string,
  city: string,
  timeout: number = 5000
) : Promise<T | TErr>
{
  Logger.Info("get_analysis", `Search: ${key} in ${city}`)
  console.log("Search: ", key, " in ", city)
  // request
  const response =  await ky(`${HOST_V1}/analysis?key=${key}&city=${city}`, {
    timeout: timeout
  })
  if (!response.ok) {
    return await response.json<TErr>()
  }
  return await response.json<T>()
}