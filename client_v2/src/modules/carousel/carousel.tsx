import React, {useState} from 'react';
import {useAnalysis} from '../../stores/analysis-store'
import classes from "./style.module.css";
import {useSelectAnalysis} from "../../stores/select-analysis-store";
import {IAnalysis, LabAndAnalysis} from "../../models/analysis";
import DialogSelectAnalysis from "./dialog-select-analysis";
import {AssertMsg} from "../../common/assert_msg";
import {useFilterStore} from "../../stores/filter-store";
import CDescription from "../../ui/description/description";

// Import Swiper styles
import {Swiper, SwiperSlide} from "swiper/react";

// Import Swiper styles
import "swiper/css";
import "swiper/css/effect-coverflow";
import "swiper/css/pagination";

// import required modules
import {EffectCoverflow, Pagination, Parallax} from "swiper";
import {Card, notification} from "antd";
import {MdOutlineFavorite, MdOutlineFavoriteBorder} from "react-icons/all";
import {useFavorite} from "../../stores/favorit-store";
import {Logger} from "../../common/logger";

function CardHeader({analysis} : {analysis: IAnalysis}) {
  const favoriteStore = useFavorite()
  const [isSelect, setIsSelect] = useState(false)

  function addToFavorite(analysis: IAnalysis) {

    if (!isSelect) {
      favoriteStore.addToFavorite(analysis)
    } else {
      favoriteStore.delete(analysis)
    }
    setIsSelect(!isSelect)

    Logger.Info("favorite", favoriteStore.selectedList)
  }

  return (
    <>
      <div
        className="flex flex-row"
      >
        <span className="grow truncate">
          { analysis.name }
        </span>
        <button
          onClick={() =>
            addToFavorite(analysis)
          }
        >
          {
            isSelect ?
              <MdOutlineFavorite
                className="w-5 h-5"
              />
              :
              <MdOutlineFavoriteBorder
                className="w-5 h-5"
              />
          }
        </button>
      </div>
    </>
  )
}

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

  function RenderSwipe(listLaboratoryTests: LabAndAnalysis) {

    function filtration(analysis: IAnalysis) {
      let inTitle = false
      if (filterStore.title.length !== 0) {
        const regex = new RegExp(filterStore.title, 'gi');
        inTitle = analysis.name.match(regex) != null;
      } else {
        inTitle = true
      }

      let inDesc = false
      if (filterStore.description.length !== 0) {
        const regex = new RegExp(filterStore.description, 'gi');
        inDesc = analysis.description.match(regex) != null;
      } else {
        inDesc = true
      }

      let inPrice = false
      const isZero = filterStore.minPrice == 0 && filterStore.maxPrice == 0
      switch (isZero) {
        case true: {
          inPrice = true
          break
        }
        case false: {
          if ((analysis.price > filterStore.minPrice && analysis.price < filterStore.maxPrice)) {
            inPrice = true
          }
          break
        }
      }

      return inTitle && inDesc && inPrice
    }

    return (
      <>
        {
          listLaboratoryTests.list?.filter((value: IAnalysis) => {
            return filtration(value)
          }).map((analysis: IAnalysis, idxAnalysis) =>
            <SwiperSlide
              key={idxAnalysis}
              className="w-72 h-72 border-2 border-main-border rounded-md select-none"
            >
              <Card
                key={idxAnalysis}
                title={<CardHeader analysis={analysis}/>}
              >
                <CDescription
                  className="h-28 truncate whitespace-normal overflow-hidden"
                >
                  {analysis.description}
                </CDescription>
                <button
                  className="bg-black text-white hover:text-black hover:bg-white hover:border-2 hover:border-black duration-300 p-2 w-full rounded-md grow-0"
                  onClick={() => openSelectAnalysis(analysis)}
                >
                  Подробнее
                </button>
              </Card>
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
                <div
                  key={idx}
                >
                    <h1
                        className={`${classes.name_lab}`}
                    >
                      {listLaboratoryTests.name_lab}
                    </h1>
                    <Swiper
                        key={idx}
                        effect={"coverflow"}
                        grabCursor={true}
                        centeredSlides={true}
                        slidesPerView={"auto"}
                        coverflowEffect={{
                          rotate: 50,
                          stretch: 0,
                          depth: 100,
                          modifier: 1,
                          slideShadows: true,
                        }}
                        pagination={true}
                        modules={[EffectCoverflow, Pagination]}
                        className="w-full pb-10"
                    >
                      {
                        RenderSwipe(listLaboratoryTests)
                      }
                    </Swiper>
                </div>
            }
          </>
        )
      }
      <DialogSelectAnalysis/>
    </div>
  );
}

export default Carousel;