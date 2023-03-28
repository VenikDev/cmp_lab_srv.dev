import React from 'react';
import classes from './description.module.css'

interface ICDescription {
  children: React.ReactNode
}

function CDescription(props: ICDescription) {
  return (
    <div className={classes.desc}>
      {props.children}
    </div>
  );
}

export default CDescription;