import React, {useState} from 'react';
import classes from './style.module.css';

function NavBar() {
  const [city, setCity] = useState("Нижний Тагил")
  const [stateDialog, setStateDialog] = useState(false)
  const nameSite = "ZдравRU"

  // Open or close dialog for select city
  const openCloseDialog = () => {
    setStateDialog(!stateDialog)
  }

  return (
    <nav className="flex">
      <ul>
        <li className={classes.name}>{ nameSite }</li>
        <li className={classes.area_visible_city}>
          <button
            onClick={() => openCloseDialog()}
            className={classes.btn_select_city}
          >
            { city }
          </button>
        </li>
      </ul>
    </nav>
  );
}

export default NavBar;