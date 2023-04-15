import {create} from "zustand";
import {IAnalysis, IListAnalysis, LabAndAnalysis} from "../models/analysis";

export const useAnalysis = create<IListAnalysis>(set => ({
  analysis: undefined,
  isLoading: false,
  addAnalysis: (newAnalysis: LabAndAnalysis) => set(
    state => ({
      analysis: newAnalysis,
    })
  ),
  changeStateLoading: () => {
    console.log("change state loading")
    set(
      state => ({
        isLoading: !state.isLoading
      })
    )
  }
}))