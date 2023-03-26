import {create} from "zustand";

interface IAnalysis {

}

interface IListAnalysis {
  analysis: IAnalysis[],
  isLoading: boolean,
  error: string,
  addAnalysis: (analysis: IAnalysis[]) => void
}

export const useAnalysis = create<IListAnalysis>(set => ({
  analysis: [],
  isLoading: false,
  error: "",
  addAnalysis: (newAnalysis: IAnalysis[]) =>
    set((state) => {
      state.analysis = [...state.analysis, ...newAnalysis]
    })
}))