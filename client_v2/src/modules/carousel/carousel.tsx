import React from 'react';
import {useAnalysis} from './store'
import CAlertError from "../../ui/alerts/error/alert-error";

interface ICarousel {

}

function Carousel(props: ICarousel) {
  const analysisIsLoading = useAnalysis(state => {
    console.log(state)
    return state.isLoading
  })

  return (
    <>
      { !analysisIsLoading ?
        <CAlertError
          className="mt-10"
        >
          Пока здесь ничего нет
        </CAlertError> : ""
      }
    </>
  );
}

export default Carousel;