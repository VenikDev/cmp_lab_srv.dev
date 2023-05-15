import {Logger} from "./logger";

export function AssertMsg<T>(condition: boolean, head: string, message: T, ) {
  const group = head.length == 0 ? "assert" : `assert/${head}`
  if (condition) {
    Logger.Error(group, message)
  }
}