import React, {FunctionComponent} from 'react';
import {FavoriteAnalysis, useFavorite} from "../../stores/favorit-store";
import FavoriteDiagram from "./modules/diagram/favoriteDiagram";
import classes from "./favorite.module.css"
import _, {groupBy} from "lodash";
import CDescription from "../../ui/description/description";
import {Collapse, Divider} from 'antd';
import {Pagination} from "swiper";
import {Swiper, SwiperSlide} from "swiper/react";
import "swiper/css";
import "swiper/css/pagination";
import {GrClose} from "react-icons/gr";

interface IFavoriteByLab {
  name: string
  listOfAnalysis: FavoriteAnalysis[]
}


function Favorite() {
  // get store
  const favoriteStore = useFavorite()

  return (
    <>
      {
        _.chain(favoriteStore.selectedList)
          .groupBy(item => item.name)
          .map((value, name) =>
            <>
              <div
                key={name}
                className={classes.f_con}
              >
                <b>
                  {name}
                </b>
              </div>
              {
                value.map((analysis, idx) =>
                  <div
                    className={classes.f_con}
                    key={idx}
                  >
                    {/* header */}
                    <div
                      className="flex flex-row"
                    >
                      <div
                        className="grow"
                      >
                        <b>{analysis.analysis.name}</b>
                      </div>
                      <button
                        onClick={() => {
                          favoriteStore.delete(analysis)
                        }}
                      >
                        <GrClose
                          className="w-5 h-5"
                        />
                      </button>
                    </div>
                    {/* body */}
                    <Divider>Описание</Divider>
                    <p>
                      {analysis.analysis.description}
                    </p>
                    <Divider>Цена</Divider>
                    <CDescription>
                      <b>
                        {analysis.analysis.price} рублей
                      </b>
                    </CDescription>
                    <a
                      className="text-center mt-2 hover:text-black"
                      href={analysis.analysis.original_url}
                      target="_blank"
                    >
                      Перейти на сайт лаборатории
                    </a>
                  </div>
                )
              }
            </>
          ).value()
      }
      <div
        className={classes.f_con}
      >
        <FavoriteDiagram/>
      </div>
    </>
  );
}

export default Favorite;