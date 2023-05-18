import React from 'react';
import classes from './strong-bold.module.css'
interface IBoldRed {
  children: React.ReactNode
}

function StrongBold(props: IBoldRed) {
  return (
    <b className={classes.strong_bold}>
      { props.children }
    </b>
  );
}

export default StrongBold;