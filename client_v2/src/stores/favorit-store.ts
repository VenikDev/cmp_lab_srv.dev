import {create} from "zustand";
import {IAnalysis} from "../models/analysis";

export type FavoriteAnalysis = {
  name: string,
  analysis: IAnalysis
}

export interface IFavoriteAnalysis {
  selectedList: FavoriteAnalysis[]
  add: (analysis: FavoriteAnalysis) => void
  delete: (analysis: FavoriteAnalysis) => void
}

export const useFavorite = create<IFavoriteAnalysis>(set => ({
  selectedList: [],
  // Add to favorites
  add: (analysis: FavoriteAnalysis) => set(state => ({
    selectedList: [...state.selectedList, analysis]
  })),
  delete(analysis: FavoriteAnalysis) {
    set(state => ({
      selectedList: state.selectedList.filter((value: FavoriteAnalysis) => {
        return value.analysis.id !== analysis.analysis.id;
      }),
    }))
  }
}))