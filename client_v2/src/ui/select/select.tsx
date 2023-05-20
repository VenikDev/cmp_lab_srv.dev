import React from 'react';
import classes from "./select.module.css";
import {Logger} from "../../common/logger";

interface ISelect {
  id: string
  list: Array<string>
  label?: string
  callbackSelect: (value: string) => void
}

const Select = (props: ISelect) => {

  const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    if (props.callbackSelect) {
      props.callbackSelect(event.target.value)
    }
  };

  return (
    <div
      className="mt-3"
    >
      {
        props.label &&
          <label
              htmlFor={props.id}
              className={classes.label}
          >
            { props.label }
          </label>
      }
      <select
        id={props.id}
        className={classes.select}
        onChange={handleSelectChange}
      >
        {
          props.list.map((element, idx) =>
            <option
              key={idx}
            >
              {element}
            </option>
          )
        }
      </select>
    </div>
  );
};

export default Select;