import {create} from "zustand";
import {IAnalysis} from "../models/analysis";

export interface IFavoriteAnalysis {
  selectedList: IAnalysis[]
  add: (analysis: IAnalysis) => void
  delete: (analysis: IAnalysis) => void
}

export const useFavorite = create<IFavoriteAnalysis>(set => ({
  selectedList: [],
  // Add to favorites
  add: (analysis: IAnalysis) => set(state => ({
    selectedList: [...state.selectedList, analysis]
  })),
  delete(analysis: IAnalysis) {
    set(state => ({
      selectedList: state.selectedList.filter((value: IAnalysis) => {
        return value.id !== analysis.id;
      }),
    }))
  }
}))