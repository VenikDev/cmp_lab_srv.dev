import React, {
  useState,
  ChangeEvent
} from 'react';
import {IInputProps} from "./models";

function CInput(props: IInputProps) {
  const [value, setValue] = useState<string>("")

  // This code declares a constant `onInput` that is a function that takes an event of type
  // `ChangeEvent<HTMLInputElement>` as an argument. When called, it sets the value of the
  // input element that triggered the event to the state value `setValue`.
  const onInput = (event: ChangeEvent<HTMLInputElement>) => {
    setValue(event.target.value)
  }

  return (
    <input
      className="w-full p-2 border-red-500 border-2 rounded-md mt-3"
      placeholder={props.placeholder}
      value={value}
      onInput={onInput}
    />
  );
}

export default CInput