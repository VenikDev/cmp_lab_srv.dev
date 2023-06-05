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
import {Drawer, notification} from "antd";
import {Link} from "react-router-dom";
import {AiOutlineMenu} from "react-icons/ai";

//  stores
import {useGlobalProperties} from "../../stores/global-properties-store";
import {useMenuStore} from "../../stores/menu-store";
import Description from "../../ui/description/description";
import {MESSAGE_TEXT, TypeNotification} from "../../common/notification/notification";


function NavBar() {
  const globalPropertiesStore = useGlobalProperties()
  const menuStore = useMenuStore()

  const [stateDialog, setStateDialog] = useState(false)
  const nameSite = "ZдравRU"

  // atn
  const [notificationApi, contextHolder] = notification.useNotification();
  const sendNotification = (desc: string) => {
    notificationApi[TypeNotification.TN_ERROR]({
      message: MESSAGE_TEXT,
      description: desc
    })
  }

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
          sendNotification("Ошибка получения города по умолчаю")
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
          sendNotification(`Ошибка получения информации об ${city}`)
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
      { contextHolder }
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
              <Link
                to="/"
                className="hover:text-black focus:text-black"
              >
                {nameSite}
              </Link>
            </div>
          </li>
          <li className={classes.area_visible_city}>
            {/*<div className="flex">*/}
            {/*  <button*/}
            {/*    className={classes.menu_btn}*/}
            {/*    onClick={menuStore.open}*/}
            {/*  >*/}
            {/*    <AiOutlineMenu*/}
            {/*      className="w-6 h-6"*/}
            {/*    />*/}
            {/*  </button>*/}
            {/*</div>*/}
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
          onClick={menuStore.close}
        >
          <Link
            to="/favorite"
            className="hover:text-black"
          >
            Открыть избранное
          </Link>
        </button>
        <Description>
          После завершения сессии, данные для сравнения анализов будут удалены
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