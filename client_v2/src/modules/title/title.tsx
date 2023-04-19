import React from 'react';
import classes from './title.module.css'

function Title() {
  return (
    <div className={classes.block_with_titles}>
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