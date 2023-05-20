import React, {useState} from 'react';
import classes from "./checkbox.module.css";
import {Simulate} from "react-dom/test-utils";
import change = Simulate.change;

interface ICheckbox {
  description?: string
  title: string
  isChecked: boolean
}

const Checkbox = (props: ICheckbox) => {
  const [isCheck, setCheck] = useState(props.isChecked)
  // const changeState = (event) => {
  //   setCheck(event.value)
  // }

  return (
    <div>
      <div
        className="flex my-2 select-none"
      >
        {/* checkbox */}
        <div className="flex items-center h-5">
          <input
            id="helper-checkbox"
            aria-describedby="helper-checkbox-text"
            type="checkbox"
            value=""
            className={classes.checkbox}
            checked={isCheck}
          />
        </div>
        {/* title */}
        <div className="ml-2 text-sm">
          <label
            htmlFor="helper-checkbox"
            className={classes.title}
          >
            {props.title}
          </label>
          {/* description */}
          {
            props.description ?
              <p
                id="helper-checkbox-text"
                className={classes.description}
              >
                {props.description}
              </p> : ""
          }
        </div>
      </div>
    </div>
  );
};

export default Checkbox;