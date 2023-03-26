import React, {useState} from 'react';
import classes from './style.module.css';
import Dialog from "../../ui/dialog/dialog";
import CRB from "../../ui/text/bold-red";
import CDescription from "../../ui/description/description";
import CInput from "../../ui/input/input";

function NavBar() {
  const [city, setCity] = useState("Нижний Тагил")
  const [stateDialog, setStateDialog] = useState(false)
  const nameSite = "ZдравRU"

  // Open or close dialog for select city
  const openCloseDialog = () => {
    console.log("state dialog for select city: ", stateDialog)
    setStateDialog(!stateDialog)
  }

  return (
    <>
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

      <Dialog
        open={stateDialog}
        callbackClose={setStateDialog}
        title="Выбрать город"
      >
        <div>
          <CInput
            value={city}
            placeholder="Введите название города"
            disabled={false}
          />
          <CDescription>
            Выберите <CRB>город</CRB>, в котором нужно искать анализы
          </CDescription>
          <div>

          </div>
        </div>
      </Dialog>
    </>
  );
}

export default NavBar;