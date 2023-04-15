import {create} from "zustand";

interface IGlobalProperties {
  isPhone: boolean,
  setIsPhone: (isPhone: boolean) => void
}

export const useGlobalProperties = create<IGlobalProperties>(set => ({
  isPhone: false,
  setIsPhone: (isPhone: boolean) => set({isPhone: isPhone})
}))