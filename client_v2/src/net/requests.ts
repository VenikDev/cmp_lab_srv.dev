import ky from "ky";
import {HOST_V1} from "./consts";

export async function getAnalysis<T>(key: string, city: string) : Promise<T> {
  console.log("Search: ", key, " in ", city)
  // request
  const response =  await ky(`${HOST_V1}/analysis?key=${key}&city=${city}`)
  return await response.json<T>()
}