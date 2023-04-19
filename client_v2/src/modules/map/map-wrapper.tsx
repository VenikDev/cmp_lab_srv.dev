import React, {useEffect, useState} from 'react';
import {load} from "@2gis/mapgl";
import {useSelectAnalysis} from "../../stores/select-analysis-store";
import {useGlobalProperties} from "../../stores/global-properties-store";
import {ICity} from "../../models/city";
import {Logger} from "../../common/logger";

interface IMapWrapper {

}

const ID_MAP_WRAPPER = "map-container"
const DEFAULT_ZOOM = 13

const Map = React.memo(
  () => {
    return (
      <div
        id={ID_MAP_WRAPPER}
        className="h-[200px] w-full">
      </div>
    );
  },
  () => true,
);

function MapWrapper(props: IMapWrapper) {
  // store
  const globalPropertiesStore = useGlobalProperties()

  // coordinates
  let lat: number = 0
  let lon: number = 0

  useEffect(() => {
    const selectCity = globalPropertiesStore.selectCity
    if (selectCity) {
      lat = Number(selectCity.coords.lat)
      lon = Number(selectCity.coords.lon)

      Logger.Info(`Coordinate of ${selectCity.name}`, [lat, lon])
    }
  }, [globalPropertiesStore.selectCity])

  useEffect(() => {
    let map: any;
    load().then((mapglAPI) => {
      map = new mapglAPI.Map(ID_MAP_WRAPPER, {
        center: [lat, lon],
        zoom: DEFAULT_ZOOM,
        key: '1abeeb58-2d6f-4c06-9842-da676355abcf',
      });
    });

    // Удаляем карту при размонтировании компонента
    return () => map && map.destroy();
  }, []);

  return (
    <Map/>
  );
}

export default MapWrapper;