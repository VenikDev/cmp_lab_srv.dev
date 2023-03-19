import React from 'react';
import classes from './dialog.module.css'

interface IDialog {
  id?: string
  open: boolean
  callbackClose: (value: boolean) => void
  children: any
}

function Dialog(props: IDialog) {
  return (
    <div className={props.open ? classes.bg : ""}>
      <dialog
        open={props.open}
        id={props.id}
        className={classes.dialog_model}
      >
        {props.children}
      </dialog>
    </div>
  );
}

export default Dialog;