import React, {
  useState,
  ChangeEvent
} from 'react';
import {IInputProps} from "./models";
import classes from "./input.module.css"

function CInput(props: IInputProps) {
  const [value, setValue] = useState<string>(props.value ? props.value : "")

  // This code declares a constant `onInput` that is a function that takes an event of type
  // `ChangeEvent<HTMLInputElement>` as an argument. When called, it sets the value of the
  // input element that triggered the event to the state value `setValue`.
  const onInput = (event: ChangeEvent<HTMLInputElement>) => {
    setValue(event.target.value)
  }

  return (
    <input
      className={classes.input_field}
      placeholder={props.placeholder}
      value={value}
      onInput={onInput}
      disabled={props.disabled ? props.disabled : true}
    />
  );
}

export default CInput