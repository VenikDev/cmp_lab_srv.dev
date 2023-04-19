import React from 'react';
import {useAnalysis} from '../../stores/analysis-store'
import classes from "./style.module.css";
import {Swiper, SwiperSlide} from "swiper/react";
import {Pagination} from "swiper";
import {useSelectAnalysis} from "../../stores/select-analysis-store";
import {IAnalysis, LabAndAnalysis} from "../../models/analysis";
import Dialog from "../../ui/dialog/dialog";
import CRB from "../../ui/text/bold-red";
import CardAnalysis from "./card-analysis";
import {Logger} from "../../common/logger";
// css
import "swiper/css";
import "swiper/css/pagination";

function Carousel() {
  // stores
  const analysisStore = useAnalysis()
  const selectAnalysisStore = useSelectAnalysis()

  function openSelectAnalysis(value: IAnalysis) {
    selectAnalysisStore.changeAnalysis(value)
    selectAnalysisStore.changeState()
  }

  function analysisEmpty() {
    return analysisStore.analysis.length == 0;
  }

  function RenderSwipe(listLaboratoryTests: LabAndAnalysis, idx: number) {
    return (
      <SwiperSlide
        key={idx}
      >
        <h1 className={classes.name_lab}>
          {listLaboratoryTests.name_lab}
        </h1>
        {
          listLaboratoryTests.list?.map((analysis: IAnalysis, idxAnalysis) =>
            <CardAnalysis
              openSelectCallback={openSelectAnalysis}
              analysis={analysis}
              key={idxAnalysis}
            />
          )
        }
      </SwiperSlide>
    )
  }

  function DialogSelectAnalysis() {
    return (
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
        <hr
          className="my-3"
        />
        <h3>
          {selectAnalysisStore.analysis?.description}
        </h3>
        <div
          className="text-right mt-2"
        >
          <h1
            className={classes.price_analysis}
          >
            {selectAnalysisStore.analysis?.price} руб.
          </h1>
        </div>
        <a
          className={classes.link_to_lab}
          href={selectAnalysisStore.analysis?.original_url}
        >
          Перейти на сайт лаборатории
        </a>
        <br/>
      </Dialog>
    )
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
        {
          analysisStore.analysis?.map((listLaboratoryTests: LabAndAnalysis, idx) =>
            listLaboratoryTests?.list.length != 0 && RenderSwipe(listLaboratoryTests, idx))
        }
      </Swiper>

      <DialogSelectAnalysis/>
    </div>
  );
}

export default Carousel;