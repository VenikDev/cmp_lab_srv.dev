import {IButtonProps} from "./models";
import classes from "./button.module.css";

export function CButton(props: IButtonProps) {
  return (
    <button
      className={classes.btn_component}
      placeholder={props.placeholder}
      disabled={props.disabled}
    >
      {props.children}
    </button>
  )
}