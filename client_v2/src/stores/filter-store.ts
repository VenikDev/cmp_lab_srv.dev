import {create} from "zustand";

type DataFilter = {
  title: string
  description: string
  price: [number, number]
}

interface IFilterStore {
  title: string
  description: string
  maxPrice: number
  initialMax: number
  initialMin: number
  setMax: (max: number) => void
  setMin: (min: number) => void
  minPrice: number
  setQuery: (query: DataFilter) => void
  // window of filter
  isOpen: boolean
  open: () => void
  close: () => void
}

export const useFilterStore = create<IFilterStore>(set => ({
  title: "",
  description: "",
  minPrice: 0,
  maxPrice: 0,
  initialMax: 0,
  initialMin: 0,
  setMax: (max: number) => set({
    maxPrice: max
  }),
  setMin: (min: number) => set({
    minPrice: min
  }),
  setQuery: (query: DataFilter) => set({
    title: query.title,
    description: query.description,
    minPrice: query.price[0],
    maxPrice: query.price[1]
  }),
  isOpen: false,
  open: () => set({
    isOpen: true
  }),
  close: () => set({
    isOpen: false
  }),
}))