import React from 'react';
import classes from './footer.module.css'

function Footer() {
  return (
    <footer
      className={classes.footer}
    >
      <div>

      </div>
      <div
        className={classes.right}
      >
        <span className="mx-auto">
          2023 г. Все права защищены.
        </span>
      </div>
    </footer>
  );
}

export default Footer;