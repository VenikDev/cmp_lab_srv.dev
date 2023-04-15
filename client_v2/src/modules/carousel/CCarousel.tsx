import React from 'react';
import {useAnalysis} from '../../stores/analysis-store'
import classes from "./modulr.module.css";
import {Swiper, SwiperSlide} from "swiper/react";
import {Pagination} from "swiper";
import Description from "../../ui/description/description";
import {truncate} from "../../common/truncate";

// css
import "swiper/css";
import "swiper/css/pagination";

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
                  className={classes.card}
                  key={index}
                >
                  <h5
                    className={classes.title}
                  >
                    {value.name}
                  </h5>
                  <p className={classes.description}>
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
                  className={classes.card}
                  key={index}
                >
                  <h5
                    className={classes.title}
                  >
                    {value.name}
                  </h5>
                  <p className={classes.description}>
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
                  className={classes.card}
                  key={index}
                >
                  <h5
                    className={classes.title}
                  >
                    {value.name}
                  </h5>
                  <p className={classes.description}>
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