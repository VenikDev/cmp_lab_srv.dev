import {create} from "zustand";
import {FiltrationTypes} from "../ui/expended-card/FiltrationTypes";

interface IFilterStore {
  query: string
  setQuery: (newQuery: string) => void
  category: string
  setCategory: (category: string) => void
}

export const useFilterStore = create<IFilterStore>(set => ({
  query: "",
  setQuery: (query: string) => set(state => ({
    query: query
  })),
  category: FiltrationTypes.SEARCH_DESCRIPTION,
  setCategory: (category: string) => set(state => ({
    query: category
  })),
}))