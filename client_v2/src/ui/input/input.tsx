import React, {
  useState,
  ChangeEvent
} from 'react';
import classes from "./input.module.css"

export interface IInputProps {
  placeholder?: string
  value?: string
  disabled?: boolean
  className?: string
  onInput?: (event: ChangeEvent<HTMLInputElement>) => void
}

function CInput(props: IInputProps) {

  return (
    <input
      className={`${classes.input_field} ${props.className}`}
      placeholder={props.placeholder}
      defaultValue={props.value}
      onChange={props.onInput}
      disabled={props.disabled ? props.disabled : false}
    />
  );
}

export default CInput