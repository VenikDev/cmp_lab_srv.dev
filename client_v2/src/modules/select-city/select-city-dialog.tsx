import React, {useEffect, useState} from 'react';
import Dialog from "../../ui/dialog/dialog";
import CDescription from "../../ui/description/description";
import CRB from "../../ui/text/bold-red";
import {useGlobalProperties} from "../../stores/global-properties-store";
import ky from "ky";
import {HOST_V1} from "../../net/consts";
import {ICity} from "../../models/city";
import classes from "./style.module.css"

interface ISelectCityDialog {
  isOpen: boolean,
  callbackClose: () => void
}

function SelectCityDialog(props: ISelectCityDialog) {
  const globalPropertiesStore = useGlobalProperties()
  const [cities, setCities] = useState<ICity[]>([])

  // get all cities for select
  useEffect(() => {
    (async () => {
      await ky(HOST_V1+`/get_list_of_cities`)
        .json<ICity[]>().then(value => {
          setCities(value)
          console.log(value)
        });
    })()
  }, [])

  function changeSelectCity(name: string) {
    globalPropertiesStore.setNewSelectCity(name)
    props.callbackClose()
  }

  // render component
  return (
    <>
      {/* select city */}
      <Dialog
        open={props.isOpen}
        callbackClose={props.callbackClose}
        title="Выбрать город"
      >
        <CDescription>
          Выберите <CRB>город</CRB>, в котором нужно искать анализы
        </CDescription>
        {/* list cities */}
        <div
          className={classes.list}
        >
          {
            cities.map((city, idx) =>
              <div
                className={classes.list_item}
                key={city.coords.lat}
                onClick={() => changeSelectCity(city.name)}
              >
                <h1
                  className={classes.list_item_text}
                >
                  {city.name}
                </h1>
                {
                  city.name == globalPropertiesStore.selectCity &&
                    <h1
                      className={classes.is_select}
                    >
                        Выбрано
                    </h1>
                }
              </div>
            )
          }
          <div>

          </div>
        </div>
      </Dialog>
    </>
  );
}

export default SelectCityDialog;