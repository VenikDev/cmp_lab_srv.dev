import {create} from "zustand";

interface IGlobalProperties {
  isPhone: boolean,
  setIsPhone: (isPhone: boolean) => void,
  selectCity: string,
  setNewSelectCity: (newCity: string) => void
}

export const useGlobalProperties = create<IGlobalProperties>(set => ({
  isPhone: false,
  setIsPhone: (isPhone: boolean) => set({isPhone: isPhone}),
  selectCity: "Нижний Тагил",
  setNewSelectCity: (newCity: string) => set({selectCity: newCity})
}))