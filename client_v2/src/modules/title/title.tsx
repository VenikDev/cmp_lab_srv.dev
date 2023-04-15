import React from 'react';
import classes from './title.module.css'

function Title() {
  return (
    <div className="mx-5">
      <div className={classes.title}>
        Инвестируй
      </div>
      <div className={classes.title + " " + classes.title_with_bg}>
        в свое
      </div>
      <div className={classes.title}>
        здоровье
      </div>
    </div>
  );
}

export default Title;