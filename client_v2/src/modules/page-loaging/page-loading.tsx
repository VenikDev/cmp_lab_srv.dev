import React, {useState} from 'react';
import Dialog from "../../ui/dialog/dialog";
import {useAnalysis} from "../../stores/analysis-store";
import waitingGif from "../../assets/waiting.gif"
import CRB from "../../ui/text/bold-red";

function PageLoading() {
  const analysisStore = useAnalysis()
  const [dots, setDots] = useState("")
  const timeout = 1000

  setInterval(() => {
    if (!analysisStore.isLoading) {
      return
    }

    if (dots.length > 5) {
      setDots("")
      return
    }

    setDots(dots + ".")
  }, timeout)

  return (
    <>
      <Dialog
        open={analysisStore.isLoading}
        title="Отправляем запрос на сервер"
        canBeClosed={false}
      >
        <img
          width="640"
          height="358"
          src={waitingGif}
          alt="Не могу прогрузить веселую гифку"
        />
        <div
          className="mt-4"
        >
          <CRB>Ждем ответа{dots}</CRB>
        </div>
      </Dialog>
    </>
  );
}

export default PageLoading;