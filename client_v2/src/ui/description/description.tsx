import React from 'react';
import classes from './description.module.css'

interface ICDescription {
  children: React.ReactNode
  className?: string
}

function CDescription(props: ICDescription) {
  return (
    <div className={`${classes.desc} ${props.className}`}>
      {props.children}
    </div>
  );
}

export default CDescription;