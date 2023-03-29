import {ChangeEvent} from "react";

export interface IInputProps {
  placeholder?: string
  value?: string
  disabled?: boolean
  className?: string
  onInput?: (event: ChangeEvent<HTMLInputElement>) => void
}