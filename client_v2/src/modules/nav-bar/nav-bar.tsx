import React, {useState} from 'react';
import classes from './style.module.css';
import Dialog from "../../ui/dialog/dialog";
import CRB from "../../ui/text/bold-red";
import CDescription from "../../ui/description/description";
import CInput from "../../ui/input/input";
import {MdFavoriteBorder} from "react-icons/all";

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
      <nav>
        <ul className="flex justify-between">
          <li className={classes.name}>{ nameSite }</li>
          <li className={classes.area_visible_city}>
            <div className="flex">
              {/* open dialog for select city */}
              <button
                onClick={() => openCloseDialog()}
                className={classes.btn_select_city}
              >
                { city }
              </button>
              {/* open page fot visible selected analysis */}
              <button
                className={classes.btn_selected}
              >
                <MdFavoriteBorder
                  className="w-5 h-5"
                />
              </button>
            </div>
          </li>
        </ul>
      </nav>

      {/* select city */}
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

      {/* selected analyzes */}

    </>
  );
}

export default NavBar;