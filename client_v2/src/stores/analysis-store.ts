import {create} from "zustand";
import {IAnalysis, IListAnalysis} from "../models/analysis";

export const useAnalysis = create<IListAnalysis>(set => ({
  analysis: [],
  isLoading: false,
  addAnalysis: (newAnalysis: IAnalysis[]) => set(
    state => ({analysis: [...state.analysis, ...newAnalysis]})
  )
}))