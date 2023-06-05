import {create} from "zustand";
import {IAnalysis} from "../models/analysis";

export interface IFavoriteAnalysis {
  selectedList: IAnalysis[]
  addToFavorite: (analysis: IAnalysis) => void
  delete: (analysis: IAnalysis) => void
}

export const useFavorite = create<IFavoriteAnalysis>(set => ({
  selectedList: [],
  // Add to favorites
  addToFavorite: (analysis: IAnalysis) => set(state => ({
    selectedList: [...state.selectedList, analysis],
    length: state.selectedList.length
  })),
  delete(analysis: IAnalysis) {
    set(state => ({
      selectedList: state.selectedList.filter((value: IAnalysis) => {
        return value.name == analysis.name;
      }),
    }))
  }
}))