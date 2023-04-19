import {create} from "zustand";
import {IAnalysis, IListAnalysis, LabAndAnalysis} from "../models/analysis";

export const useAnalysis = create<IListAnalysis>(set => ({
  analysis: [],
  isLoading: false,
  addAnalysis: (newAnalysis: LabAndAnalysis[]) => {
    set({
      analysis: newAnalysis
    })
  },
  changeStateLoading: () => {
    set(
      state => ({
        isLoading: !state.isLoading
      })
    )
  }
}))