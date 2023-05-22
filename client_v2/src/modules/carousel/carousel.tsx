import React from 'react';
import {useAnalysis} from '../../stores/analysis-store'
import classes from "./style.module.css";
import {useSelectAnalysis} from "../../stores/select-analysis-store";
import {IAnalysis, LabAndAnalysis} from "../../models/analysis";
import CardAnalysis from "./card-analysis";
import DialogSelectAnalysis from "./dialog-select-analysis";
import {AssertMsg} from "../../common/assert_msg";
import {Logger} from "../../common/logger";
import {FiltrationTypes} from "../../ui/expended-card/FiltrationTypes";
import {useFilterStore} from "../../stores/filter-store";

// Import Swiper styles
import {Swiper, SwiperSlide} from "swiper/react";

// Import Swiper styles
import "swiper/css";
import "swiper/css/effect-coverflow";
import "swiper/css/pagination";

// import required modules
import {EffectCoverflow, Pagination, Parallax} from "swiper";

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

  function RenderSwipe(listLaboratoryTests: LabAndAnalysis) {
    return (
      <>
        {
          listLaboratoryTests.list?.filter((value: IAnalysis) => {
            if (filterStore.query.length !== 0) {
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
            }
            return true
          }).map((analysis: IAnalysis, idxAnalysis) =>
            <SwiperSlide
              className="w-[300px] h-[300px]"
            >
              <CardAnalysis
                key={idxAnalysis}
                openSelectCallback={openSelectAnalysis}
                analysis={analysis}
              />
            </SwiperSlide>
          )
        }
      </>
    )
  }

  return (
    <div>
      {
        analysisStore.analysis?.map((listLaboratoryTests: LabAndAnalysis, idx) =>
          <>
            {
              listLaboratoryTests?.list.length != 0 &&
                <>
                    <h1
                        className={`${classes.name_lab}`}
                    >
                      {listLaboratoryTests.name_lab}
                    </h1>
                    <Swiper
                        effect={"coverflow"}
                        grabCursor={true}
                        centeredSlides={true}
                        slidesPerView={"auto"}
                        coverflowEffect={{
                          rotate: 50,
                          stretch: 0,
                          depth: 100,
                          modifier: 1,
                          // slideShadows: true,
                        }}
                        pagination={true}
                        modules={[EffectCoverflow, Pagination]}
                        className="w-full"
                    >
                      {
                        RenderSwipe(listLaboratoryTests)
                      }
                    </Swiper>
                </>
            }
          </>
        )
      }
      <DialogSelectAnalysis/>
    </div>
  );
}

export default Carousel;