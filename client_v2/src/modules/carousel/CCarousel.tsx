import React from 'react';
import {useAnalysis} from '../../stores/analysis-store'
import CInfoAlert from "../../ui/alerts/info/info-alert";

interface ICarousel {

}

function CCarousel(props: ICarousel) {
  const analysisIsLoading = useAnalysis()

  return (
    <>
      { !analysisIsLoading.isLoading ?
        <CInfoAlert
          className="mt-10 max-w-2xl p-5 mx-auto"
        >
          Пока здесь ничего нет
        </CInfoAlert> : ""
      }
    </>
  );
}

export default CCarousel;