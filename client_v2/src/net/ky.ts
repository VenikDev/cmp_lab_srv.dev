import ky from "ky";
import {HOST_V1} from "./consts";


export const api = ky.create({
  prefixUrl: HOST_V1
})