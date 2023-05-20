import React, {useEffect, useState} from 'react';
import classes from './style.module.css';
import SelectCityDialog from "../select-city/select-city-dialog";
import {useGlobalProperties} from "../../stores/global-properties-store";
import ky from "ky";
import {HOST_V1} from "../../net/consts";
import {ICity} from "../../models/city";
import {Logger} from "../../common/logger";
import {KEY_SELECT_CITY_FOR_LOCAL_STORE} from "../../common/keys";
import {IError} from "../../models/error";


function NavBar() {
  const globalPropertiesStore = useGlobalProperties()

  const [stateDialog, setStateDialog] = useState(false)
  const nameSite = "ZдравRU"

  useEffect(() => {
    let city = window.localStorage.getItem(KEY_SELECT_CITY_FOR_LOCAL_STORE)
    Logger.Info("city from local store", city)

    if (city == null) {
      (async () => {
        let response = await ky(HOST_V1+"/get_default_city")
        if (response.ok) {
          const defaultCity = await response.json<ICity>()

          globalPropertiesStore.setNewSelectCity(defaultCity)
        }
        else {
          const error = await response.json<IError>()
          Logger.Error("get city", error.message)
        }
      })()
    } else {
      (async () => {
        let response = await ky(HOST_V1+`/get_city_info?city=${city}`)
        if (response.ok) {
          const cityInfo = await response.json<ICity>()

          globalPropertiesStore.setNewSelectCity(cityInfo)
        }
        else {
          const error = await response.json<IError>()
          Logger.Error("get city", error.message)
        }
      })()
    }

    // return () => setTimeout(() => {Logger.Info("nav bar", "unsubscribe useEffect")}, 1000)
  }, [])

  // Open or close dialog for select city
  const openCloseDialog = () => {
    console.log("state dialog for select city: ", stateDialog)
    setStateDialog(!stateDialog)
  }

  return (
    <>
      <nav>
        <ul className="flex justify-between">
          <li>
            <div
              className={classes.name}
            >
              {nameSite}
            </div>
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
                  {globalPropertiesStore.selectCity?.name ?? "Выберите город"}
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