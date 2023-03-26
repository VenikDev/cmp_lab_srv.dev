import React, {useEffect, useState} from 'react';
import Dialog from "../../ui/dialog/dialog";
import {useAnalysis} from "../../stores/analysis-store";

function PageLoading() {
  const [stateLoading, useStateLoading] = useState(useAnalysis(state => {
    console.log("Loading state: ", state.isLoading)
    return state.isLoading
  }))

  const [dots, setDots] = useState("")
  useEffect(() => {
    setInterval(() => {
      if (dots.length  >= 3) {
        setDots("")
        return
      }
      setDots(dots + ".")
    }, 2000)
  }, [])

  return (
    <>
      <Dialog
        open={stateLoading}
        callbackClose={useStateLoading}
        title="Отправляем запрос на сервер"
      >
        { dots }
      </Dialog>
    </>
  );
}

export default PageLoading;