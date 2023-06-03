import React, {useEffect, useState} from 'react';
import {IPopular} from "./model";
import ky from "ky";
import {HOST_V1} from "../../net/consts";
import CDescription from "../../ui/description/description";
import CRB from "../../ui/text/strong_bold";
import {getAnalysis} from "../../net/requests";
import {Swiper, SwiperSlide} from "swiper/react";
import {Pagination} from "swiper";
import {LabAndAnalysis} from "../../models/analysis";
import {useAnalysis} from "../../stores/analysis-store";
import {useGlobalProperties} from "../../stores/global-properties-store";
import {Logger} from "../../common/logger";
import {IError} from "../../models/error";

// css
import classes from "./popular.module.css"
import "swiper/css";
import "swiper/css/pagination";
import {message, notification} from "antd";
import {minMaxFilter} from "../../common/minMaxFilter";
import {useFilterStore} from "../../stores/filter-store";

function Popular() {
  const globalPropertiesStore = useGlobalProperties()
  const analysisStore = useAnalysis()
  const filterStore = useFilterStore()

  const [popular, setPopular] = useState<IPopular[]>()

  // ant
  const [notificationApi, contextHolder] = notification.useNotification();
  // To check if the variable popular is null or undefined.
  // If popular is not null or undefined, then it checks the length of the array using the length
  // property and returns a boolean value indicating whether the length is equal to zero or not.
  function popularEmpty() {
    return popular?.length === 0
  }

  // await ky(HOST_V1 + "/get_popular"): This line sends a GET request to a Redis server at the "/get_popular"
  // endpoint using the ky library. HOST_V1 is a constant that represents the base URL of the Redis service.
  //
  // .json<IPopular[]>(): This line parses the response as JSON and casts it as an array of IPopular
  // objects (as defined elsewhere in the code).
  //
  // .then(value => setPopular(value)): This line sets the state of the component to the array
  // of IPopular objects returned by the Redis server. This will trigger a re-render of the component.
  useEffect(() => {
    (async () => {
      // get popular analysis from redis
      try {
        const response = await ky(HOST_V1 + "/get_popular")
        // response is fail
        if (response.status !== 200) {
          const errResponse = await response.json<{error: string}>()
          Logger.Warring("req/get_popular", errResponse.error)

          notificationApi['error']({
            message: "Ошибка получения популярного",
            description: errResponse.error
          })
        } else {
          response.json<IPopular[]>().then(value => setPopular(value))
        }
      } catch (ex) {
        notificationApi['error']({
          message: "Ошибка получения популярного"
        })
      }
    })()
  }, [])

  function sendNotification(msg: string) {
    notificationApi['error']({
      message: "Ошибка",
      description: msg
    })
  }

  // render component
  return (
    !popularEmpty() ? <div
      className="w-full my-4"
    >
      {contextHolder}
      <Swiper
        slidesPerView={3}
        spaceBetween={10}
        pagination={{
          clickable: true,
        }}
        breakpoints={{
          "@0.00": {
            slidesPerView: 3,
            spaceBetween: 10,
          },
          "@0.75": {
            slidesPerView: 4,
            spaceBetween: 20,
          },
          "@1.00": {
            slidesPerView: 5,
            spaceBetween: 40,
          },
          "@1.50": {
            slidesPerView: 6,
            spaceBetween: 50,
          },
        }}
        modules={[Pagination]}
        className="p-4"
      >
        {
          popular && popular.length != 0 && popular.map((item, id) =>
            <SwiperSlide
              key={id}
              className={classes.slide}
            >
              <div
                className="cursor-pointer"
                onClick={async () => {
                  analysisStore.changeStateLoading()
                  const analysisOrErr
                    = await getAnalysis<Array<LabAndAnalysis>>(item.name, globalPropertiesStore.selectCity?.name!!)

                  let messageErr = (analysisOrErr as IError).message
                  if (messageErr !== undefined) {
                    Logger.Warring("popular/select_city", analysisOrErr)
                    sendNotification(messageErr)
                  } else {
                    const labs = analysisOrErr as Array<LabAndAnalysis>
                    const [min, max] = minMaxFilter(labs)

                    filterStore.setMin(min)
                    filterStore.setMax(max)

                    analysisStore.addAnalysis(labs)
                  }

                  analysisStore.changeStateLoading()
                }}
              >
                <CRB>
                  {item.name}
                </CRB>
                {/* count */}
                <div
                  className="text-center"
                >
                  Искали <CRB>{item.count}</CRB> раз
                </div>

                <CDescription
                  className={classes.description}
                >
                  Нажните, чтобы искать
                </CDescription>
              </div>
            </SwiperSlide>
          )
        }
      </Swiper>
      {
        popular && popular.length != 0 &&
        <CDescription
            className="text-center"
        >
            Потяните <CRB>влево</CRB> или <CRB>вправо</CRB>, чтоб посмотреть еще
        </CDescription>
      }
    </div> : <></>
  );
}

export default Popular;