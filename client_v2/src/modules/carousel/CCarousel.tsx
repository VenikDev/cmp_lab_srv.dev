import React from 'react';
import {useAnalysis} from '../../stores/analysis-store'
import classes from "./modulr.module.css";
import {IAnalysis} from "../../models/analysis";
import {Swiper, SwiperSlide} from "swiper/react";
import {Pagination} from "swiper";

// css
import "swiper/css";
import "swiper/css/pagination";
import Description from "../../ui/description/description";
import {truncate} from "../../common/truncate";

function CCarousel() {
  const analysisStore = useAnalysis()

  return (
    <div>
      <Swiper
        slidesPerView={3}
        spaceBetween={10}
        pagination={{
          clickable: true,
        }}
        breakpoints={{
          "@0.00": {
            slidesPerView: 1,
            spaceBetween: 10,
          },
          "@0.50": {
            slidesPerView: 2,
            spaceBetween: 20,
          },
          "@1.00": {
            slidesPerView: 3,
            spaceBetween: 40,
          },
          "@1.50": {
            slidesPerView: 3,
            spaceBetween: 50,
          },
        }}
        modules={[Pagination]}
        className="p-4"
      >
        {analysisStore.analysis ?
          <div>
            <SwiperSlide>
              <h1>Citilab</h1>

              {analysisStore.analysis.citilab?.map((value, index) =>
                <div
                  // onClick={}
                  className="block p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 mb-4 cursor-pointer"
                  key={index}
                >
                  <h5 className="mb-2 text-2xl font-bold tracking-tight text-red-500">
                    {value.name}
                  </h5>
                  <p className="font-normal text-gray-700 dark:text-gray-400">
                    {truncate(value.description)}
                  </p>
                  <Description
                    className="text-center mt-2"
                  >
                    Нажми, чтобы узнать подробнее
                  </Description>
                </div>
              )}
            </SwiperSlide>
            <SwiperSlide>
              <h1>Gemotest</h1>

              {analysisStore.analysis.gemotest?.map((value, index) =>
                <div
                  // onClick={}
                  className="block p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 mb-4 cursor-pointer"
                  key={index}
                >
                  <h5 className="mb-2 text-2xl font-bold tracking-tight text-red-500">
                    {value.name}
                  </h5>
                  <p className="font-normal text-gray-700 dark:text-gray-400">
                    {truncate(value.description)}
                  </p>
                  <Description
                    className="text-center mt-2"
                  >
                    Нажми, чтобы узнать подробнее
                  </Description>
                </div>
              )}
            </SwiperSlide>
            <SwiperSlide>
              <h1>Invitro</h1>

              {analysisStore.analysis.invitro?.map((value, index) =>
                <div
                  // onClick={}
                  className="block p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 mb-4 cursor-pointer"
                  key={index}
                >
                  <h5 className="mb-2 text-2xl font-bold tracking-tight text-red-500">
                    {value.name}
                  </h5>
                  <p className="font-normal text-gray-700 dark:text-gray-400">
                    {truncate(value.description)}
                  </p>
                  <Description
                    className="text-center mt-2"
                  >
                    Нажми, чтобы узнать подробнее
                  </Description>
                </div>
              )}
            </SwiperSlide>
          </div> : ""
        }
      </Swiper>
    </div>
  );
}

export default CCarousel;