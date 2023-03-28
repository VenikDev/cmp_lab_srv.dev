import React from "react";

export interface IDialog {
  id?: string
  open: boolean
  callbackClose?: (value: boolean) => void
  children: React.ReactNode
  title?: string
  className?: string
  canBeClosed?: boolean
}