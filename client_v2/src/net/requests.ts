import ky from "ky";
import {HOST_V1} from "./consts";
import {Logger} from "../common/logger";
import {IError} from "../models/error";

export async function getAnalysis<T, TErr = IError>(
  key: string,
  city: string,
  timeout: number = 10000 // 10 сек
) : Promise<T | TErr>
{
  Logger.Info("get_analysis", `Search: ${key} in ${city}`)
  console.log("Search: ", key, " in ", city)

  // request
  const response =  await ky.get(`${HOST_V1}/analysis?key=${key}&city=${city}`)

  if (!response.ok) {
    const error = await response.json<TErr>()
    Logger.Error("get_analysis", error)
    return error
  }
  return await response.json<T>()
}