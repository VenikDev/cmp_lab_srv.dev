import {create} from "zustand";

interface IFilterStore {
  query: string | null
  newQuery: (newQuery: string) => void
}

export const useFavorite = create<IFilterStore>(set => ({
  query: null,
  newQuery: (newQuery: string) => set(state => ({
    query: newQuery
  }))
}))