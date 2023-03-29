import {create} from "zustand";

interface IStoreCity {
  city: string
  setCity: (name: string) => void
}

export const useCityStore = create<IStoreCity>(set => ({
  // default city
  city: "Нижний тагил",
  // set new city
  setCity: (name: string) => {
    console.log("change target city")
    set(
      state => ({
        city: name
      })
    )
  }
}))