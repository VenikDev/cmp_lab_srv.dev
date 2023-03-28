import {create} from "zustand";
import {IAnalysis, IFavoriteAnalysis} from "../models/analysis";

export const useFavorite = create<IFavoriteAnalysis>(set => ({
  list: [],
  length: 0,
  // Add to favorites
  addToFavorite: (analysis: IAnalysis) => set(state => ({
    list: [...state.list, analysis],
    length: state.list.length
  }))
}))