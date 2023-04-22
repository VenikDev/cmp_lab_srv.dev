import {Logger} from "./logger";

export function assert_msg<T>(condition: boolean, msg: T) {
  if (condition) {
    Logger.Error("assert", msg)
  }
}