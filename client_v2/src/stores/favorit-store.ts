import {create} from "zustand";
import {IAnalysis} from "../models/analysis";

export interface IFavoriteAnalysis {
  list: IAnalysis[]
  addToFavorite: (analysis: IAnalysis) => void
}

export const useFavorite = create<IFavoriteAnalysis>(set => ({
  list: [],
  // Add to favorites
  addToFavorite: (analysis: IAnalysis) => set(state => ({
    list: [...state.list, analysis],
    length: state.list.length
  }))
}))