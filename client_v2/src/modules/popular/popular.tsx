import React, {useEffect, useState} from 'react';
import {IPopular} from "./model";
import ky from "ky";
import {HOST_V1} from "../../net/consts";
import CDescription from "../../ui/description/description";
import CRB from "../../ui/text/bold-red";
import {getAnalysis} from "../../net/requests";
import {Swiper, SwiperSlide} from "swiper/react";
import {Pagination} from "swiper";

// css
import classes from "./popular.module.css"
import "swiper/css";
import "swiper/css/pagination";

function Popular() {
  const [popular, setPopular] = useState<IPopular[]>()

  useEffect(() => {
    const getPopular = async () => {
      // get popular analysis from redis
      await ky(HOST_V1 + "/get_popular")
        .json<IPopular[]>()
        .then(value => setPopular(value))
    }

    getPopular()
    console.log(popular)
  }, [])

  return (
    <div
      className="w-full my-4"
    >
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
          popular && popular.map((item, id) =>
            <SwiperSlide
              key={id}
              className="border-2 border-red-500 rounded-md text-center"
            >
              <div
                className="cursor-pointer"
                onClick={() => getAnalysis(item.name, "")}
              >
                <CRB>{item.name}</CRB>
                {/* count */}
                <div
                  className="text-center"
                >
                  Искали <CRB>{item.count}</CRB> раз
                </div>

                <CDescription
                  className="text-xs text-center"
                >
                  Нажните, чтобы искать
                </CDescription>
              </div>
            </SwiperSlide>
          )
        }
      </Swiper>
    </div>
  );
}

export default Popular;