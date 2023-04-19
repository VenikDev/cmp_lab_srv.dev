import {create} from "zustand";

const KEY_SELECT_CITY_FOR_LOCAL_STORE = "zdravru_select_city"

interface IGlobalProperties {
  isPhone: boolean,
  setIsPhone: (isPhone: boolean) => void,
  selectCity: string,
  setNewSelectCity: (newCity: string) => void
}

export const useGlobalProperties = create<IGlobalProperties>(set => ({
  isPhone: false,
  setIsPhone: (isPhone: boolean) => set({isPhone: isPhone}),
  selectCity: window.localStorage.getItem(KEY_SELECT_CITY_FOR_LOCAL_STORE) ?? "Нижний Тагил",
  setNewSelectCity: (newCity: string) => {
    window.localStorage.setItem(KEY_SELECT_CITY_FOR_LOCAL_STORE, newCity)
    set({selectCity: newCity})
  }
}))