import React from 'react';
import {useAnalysis} from '../../stores/analysis-store'
import classes from "./style.module.css";
import {Swiper, SwiperSlide} from "swiper/react";
import {Pagination} from "swiper";
import {useSelectAnalysis} from "../../stores/select-analysis-store";
import {IAnalysis, LabAndAnalysis} from "../../models/analysis";
import CardAnalysis from "./card-analysis";
import DialogSelectAnalysis from "./dialog-select-analysis";
import {assert_msg} from "../../common/assert_msg";

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
    const condition = analysisStore.analysis.length == 0
    assert_msg(condition, analysisStore.analysis)
    return analysisStore.analysis.length == 0
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