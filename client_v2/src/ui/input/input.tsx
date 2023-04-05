import React, {
  useState,
  ChangeEvent
} from 'react';
import {IInputProps} from "./models";
import classes from "./input.module.css"

function CInput(props: IInputProps) {

  return (
    <input
      className={`${classes.input_field} ${props.className}`}
      placeholder={props.placeholder}
      value={props.value}
      onChange={props.onInput}
      disabled={props.disabled ? props.disabled : false}
    />
  );
}

export default CInput