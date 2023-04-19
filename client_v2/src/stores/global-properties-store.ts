import {create} from "zustand";
import {ICity} from "../models/city";
import {KEY_SELECT_CITY_FOR_LOCAL_STORE} from "../common/keys";
import {Logger} from "../common/logger";

interface IGlobalProperties {
  isPhone: boolean,
  setIsPhone: (isPhone: boolean) => void,
  selectCity: ICity | undefined,
  setNewSelectCity: (newCity: ICity) => void
}

export const useGlobalProperties = create<IGlobalProperties>(set => ({
  isPhone: false,
  setIsPhone: (isPhone: boolean) => set({isPhone: isPhone}),
  selectCity: undefined,
  setNewSelectCity: (newCity: ICity) => {
    Logger.Info("select city store", newCity)

    window.localStorage.setItem(KEY_SELECT_CITY_FOR_LOCAL_STORE, newCity.name)
    set({selectCity: newCity})
  }
}))