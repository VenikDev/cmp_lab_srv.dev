import React, {
  useState,
  ChangeEvent
} from 'react';
import {IInputProps} from "./models";

export function CInput(props: IInputProps) {
  const [value, setValue] = useState<string>("")

  // This code declares a constant `onInput` that is a function that takes an event of type
  // `ChangeEvent<HTMLInputElement>` as an argument. When called, it sets the value of the
  // input element that triggered the event to the state value `setValue`.
  const onInput = (event: ChangeEvent<HTMLInputElement>) => {
    setValue(event.target.value)
  }

  return (
    <input
      placeholder={props.placeholder}
      value={value}
      onInput={onInput}
    />
  );
}