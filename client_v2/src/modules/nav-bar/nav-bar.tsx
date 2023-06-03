import React, {useEffect, useState} from 'react';
import classes from './style.module.css';
import btn_class from '../../ui/btn.module.css'
import SelectCityDialog from "../select-city/select-city-dialog";
import ky from "ky";
import {HOST_V1} from "../../net/consts";
import {ICity} from "../../models/city";
import {Logger} from "../../common/logger";
import {KEY_SELECT_CITY_FOR_LOCAL_STORE} from "../../common/keys";
import {IError} from "../../models/error";
import {AiOutlineMenu} from "react-icons/all";
import {Drawer} from "antd";

//  stores
import {useGlobalProperties} from "../../stores/global-properties-store";
import {useMenuStore} from "../../stores/menu-store";
import Description from "../../ui/description/description";


function NavBar() {
  const globalPropertiesStore = useGlobalProperties()
  const menuStore = useMenuStore()

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
      <div
        className={classes.city}
      >
        <button
          onClick={() => openCloseDialog()}
          className={classes.btn_select_city}
        >
          Город:&nbsp;
          <b className="underline-offset-2 underline">
            {globalPropertiesStore.selectCity?.name ?? "Выберите город"}
          </b>
        </button>
      </div>
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
              <button
                className={classes.menu_btn}
                onClick={menuStore.open}
              >
                <AiOutlineMenu/>
              </button>
            </div>
          </li>
        </ul>
      </nav>

      <Drawer
        title="Меню"
        placement={menuStore.placement}
        onClose={menuStore.close}
        open={menuStore.isOpen}
      >
        <button
          className={btn_class.btn}
        >
          Открыть избранное
        </button>
        <Description>
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Commodi deleniti doloribus qui quibusdam tempora. Aut commodi, dolore dolorum, eligendi impedit ipsa ipsum, nostrum odio possimus quia reiciendis sapiente suscipit ut.
        </Description>
      </Drawer>

      <SelectCityDialog
        isOpen={stateDialog}
        callbackClose={openCloseDialog}
      />

    </>
  );
}

export default NavBar;