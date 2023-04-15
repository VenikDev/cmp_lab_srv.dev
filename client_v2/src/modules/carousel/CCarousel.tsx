import React from 'react';
import {useAnalysis} from '../../stores/analysis-store'
import classes from "./modulr.module.css";
import {Swiper, SwiperSlide} from "swiper/react";
import {Pagination} from "swiper";
import Description from "../../ui/description/description";
import {truncate} from "../../common/truncate";
import {useSelectAnalysis} from "../../stores/select-analysis-store";
import {IAnalysis} from "../../models/analysis";
import Dialog from "../../ui/dialog/dialog";
import CRB from "../../ui/text/bold-red";

// css
import "swiper/css";
import "swiper/css/pagination";


function CCarousel() {
  // stores
  const analysisStore = useAnalysis()
  const selectAnalysisStore = useSelectAnalysis()

  function openSelectAnalysis(value: IAnalysis) {
    selectAnalysisStore.changeAnalysis(value)
    selectAnalysisStore.changeState()
  }

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
              <h1 className={classes.name_lab}>
                Citilab
              </h1>

              {analysisStore.analysis.citilab?.map((value, index) =>
                <div
                  onClick={() => openSelectAnalysis(value)}
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
              <h1 className={classes.name_lab}>
                Gemotest
              </h1>

              {analysisStore.analysis.gemotest?.map((value, index) =>
                <div
                  onClick={() => openSelectAnalysis(value)}
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
              <h1 className={classes.name_lab}>
                Invitro
              </h1>

              {analysisStore.analysis.invitro?.map((value, index) =>
                <div
                  onClick={() => openSelectAnalysis(value)}
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

      <Dialog
        open={selectAnalysisStore.isOpen}
        title="Подробности анализа"
        callbackClose={selectAnalysisStore.changeState}
      >
        <h1>
          <CRB>
            {selectAnalysisStore.analysis?.name}
          </CRB>
        </h1>
        <h3>
          {selectAnalysisStore.analysis?.description}
        </h3>
        <div
          className="text-right mt-2"
        >
          <CRB>
            Цена: {selectAnalysisStore.analysis?.price}
          </CRB>
        </div>
        <a
          className="text-red-500 text-center hover:text-red-900 duration-200 font-bold mt-4"
          href={selectAnalysisStore.analysis?.original_url}
        >
          Перейти на сайт лаборатории
        </a>
      </Dialog>
    </div>
  );
}

export default CCarousel;