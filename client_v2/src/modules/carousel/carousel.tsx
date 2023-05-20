import React from 'react';
import {useAnalysis} from '../../stores/analysis-store'
import classes from "./style.module.css";
import {Swiper, SwiperSlide} from "swiper/react";
import {Pagination} from "swiper";
import {useSelectAnalysis} from "../../stores/select-analysis-store";
import {IAnalysis, LabAndAnalysis} from "../../models/analysis";
import CardAnalysis from "./card-analysis";
import DialogSelectAnalysis from "./dialog-select-analysis";
import {AssertMsg} from "../../common/assert_msg";
import {Logger} from "../../common/logger";

// css
import "swiper/css";
import "swiper/css/pagination";
import {useFilterStore} from "../../stores/filter-store";
import category from "../../ui/category/category";
import {FiltrationTypes} from "../../ui/expended-card/FiltrationTypes";

function Carousel() {
  // stores
  const analysisStore = useAnalysis()
  const selectAnalysisStore = useSelectAnalysis()
  const filterStore = useFilterStore()

  function openSelectAnalysis(value: IAnalysis) {
    selectAnalysisStore.changeAnalysis(value)
    selectAnalysisStore.changeState()
  }

  function analysisEmpty() {
    const condition = analysisStore.analysis.length == 0
    AssertMsg(condition, "carousel", analysisStore.analysis)

    return analysisStore.analysis.length == 0
  }

  function getStyleByNameLab(name: string, tag: string): string | undefined {
    switch (name) {
      case "gemotest": {
        return tag + "-gemotest"
      }
      case "citilab": {
        return tag + "-citilab"
      }
      case "invitro": {
        return tag + "-invitro"
      }
    }
  }

  function RenderSwipe(listLaboratoryTests: LabAndAnalysis, idx: number) {
    const color = getStyleByNameLab(listLaboratoryTests.name_lab, "text")!!
    Logger.Info("color", color)

    return (
      <SwiperSlide
        key={idx}
      >
        <h1 className={`${classes.name_lab} ${color}`}>
          {listLaboratoryTests.name_lab}
        </h1>
        {
          listLaboratoryTests.list?.filter((value: IAnalysis) => {
            const regex = new RegExp(filterStore.query, 'gi');

            const category = filterStore.category
            switch (category) {
              case FiltrationTypes.SEARCH_DESCRIPTION: {
                return value.description.match(regex) != null;
              }
              default: {
                return value.name.match(regex) != null;
              }
            }

          }).map((analysis: IAnalysis, idxAnalysis) =>
            <CardAnalysis
              key={idxAnalysis}
              openSelectCallback={openSelectAnalysis}
              analysis={analysis}
              colorLab={color}
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
            spaceBetween: 20,
          },
          "@0.75": {
            slidesPerView: 2,
            spaceBetween: 20,
          },
          "@1.50": {
            slidesPerView: 2,
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