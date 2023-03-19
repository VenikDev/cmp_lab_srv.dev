import {IButtonProps} from "./models";
import classes from "./button.module.css";

export function CButton(props: IButtonProps) {
  const styles = classes.btn_component + " comfortaa " + props.className;
  return (
    <button
      className={styles}
      placeholder={props.placeholder}
      disabled={props.disabled}
    >
      {props.children}
    </button>
  )
}