import React, {useState} from 'react';
import classes from './style.module.css';
import Dialog from "../../ui/dialog/dialog";
import CRB from "../../ui/text/bold-red";
import CDescription from "../../ui/description/description";
import CInput from "../../ui/input/input";
import {MdFavoriteBorder} from "react-icons/all";
import {Link} from "react-router-dom";
import SelectCityDialog from "../select-city/select-city-dialog";
import {useGlobalProperties} from "../../stores/global-properties-store";

function NavBar() {
  const globalPropertiesStore = useGlobalProperties()

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
          <li className={classes.name}>
            <Link to="/">
              {nameSite}
            </Link>
          </li>
          <li className={classes.area_visible_city}>
            <div className="flex">
              {/* open dialog for select city */}

              <button
                onClick={() => openCloseDialog()}
                className={classes.btn_select_city}
              >
                Город:&nbsp;
                <b className="underline-offset-2 underline">
                  {globalPropertiesStore.selectCity}
                </b>
              </button>
              {/* open page fot visible selected analysis */}
              {/*<Link to="/favorite">*/}
              {/*  <button*/}
              {/*    className={classes.btn_selected}*/}
              {/*  >*/}
              {/*    Изб*/}
              {/*    <MdFavoriteBorder*/}
              {/*      className="w-5 h-5 ml-1"*/}
              {/*    />*/}
              {/*  </button>*/}
              {/*</Link>*/}
            </div>
          </li>
        </ul>
      </nav>

      <SelectCityDialog
        isOpen={stateDialog}
        callbackClose={openCloseDialog}
      />

    </>
  );
}

export default NavBar;