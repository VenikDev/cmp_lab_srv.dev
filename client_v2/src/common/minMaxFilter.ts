import {findMaxWithPred, findMinWithPred} from "./algo/max";
import {IAnalysis, LabAndAnalysis} from "../models/analysis";
import {Logger} from "./logger";

export function minMaxFilter(labs: Array<LabAndAnalysis>): [number, number] {
  let max = 0
  let min = 0

  for (let lab of labs) {
    const maxLab = findMaxWithPred(
      lab.list,
      (lab: IAnalysis) => {
        return lab.price > 0
      })
    if (maxLab) {
      if (maxLab.price > max) {
        max = maxLab.price
      }
    }

    const minLab = findMinWithPred(
      lab.list,
      (lab: IAnalysis) => {
        return lab.price > 0
      })
    if (minLab) {
      if (minLab.price < min) {
        min = minLab.price
      }
    }
  }

  Logger.Info("useAnalysis/new", `max: ${max}`)
  Logger.Info("useAnalysis/new", `min: ${min}`)

  return [min, max]
}