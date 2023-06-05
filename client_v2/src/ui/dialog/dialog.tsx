import React from 'react';
import classes from './dialog.module.css'
import {TfiClose} from "react-icons/tfi";
import {IDialog} from "./model";

function Dialog(props: IDialog) {
  return (
    <div className={props.open ? classes.bg : ""}>
      <dialog
        open={props.open}
        id={props.id}
        className={`${classes.dialog_model} ${props.className}`}
      >
        <div className="flex justify-around">
          <h1 className="block font-bold grow">
            { props.title }
          </h1>
          {
            props.canBeClosed || props.canBeClosed == undefined ?
              <button
                className="cursor-pointer"
                onClick={() => props.callbackClose && props.callbackClose(false)}
              >
                <TfiClose/>
              </button>
            : ""
          }
        </div>
        <hr className="my-2"/>
        {props.children}
      </dialog>
    </div>
  );
}

export default Dialog;