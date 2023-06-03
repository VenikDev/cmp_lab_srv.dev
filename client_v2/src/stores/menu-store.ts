import {create} from "zustand";
import {Placement, PlacementMenu} from "./placement";

interface IMenuStore {
  isOpen: boolean
  placement: PlacementMenu
  open: () => void
  close: () => void
  setPlacement: (placement: PlacementMenu) => void
}

export const useMenuStore = create<IMenuStore>(
  set => ({
    isOpen: false,
    placement: Placement.PLACEMENT_RIGHT,
    open: () => set({
      isOpen: true
    }),
    close: () => set({
      isOpen: false
    }),
    setPlacement: (placement: PlacementMenu) => set({
      placement: placement
    })
  })
)