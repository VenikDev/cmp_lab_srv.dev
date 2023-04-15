import {create} from "zustand";
import {IAnalysis} from "../models/analysis";

interface ISelectAnalysis {
  analysis: IAnalysis | null,
  isOpen: boolean,
  changeState: () => void
  changeAnalysis: (newAnalysis: IAnalysis) => void
}

export const useSelectAnalysis = create<ISelectAnalysis>(set => ({
  analysis: null,
  isOpen: false,
  changeState: () => set(state => ({
    isOpen: !state.isOpen
  })),
  changeAnalysis: (newAnalysis: IAnalysis) => set(state => ({
    analysis: newAnalysis
  }))
}))