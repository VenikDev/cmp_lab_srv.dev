import React, {useState} from 'react';
import Dialog from "../../ui/dialog/dialog";
import {useAnalysis} from "../../stores/analysis-store";
import waitingGif1 from "../../assets/waiting.gif"
import CRB from "../../ui/text/strong_bold";

function PageLoading() {
  const analysisStore = useAnalysis()

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
          src={waitingGif1}
          alt="Не могу прогрузить веселую гифку"
          className="justify-content-center"
        />
        <div
          className="mt-4"
        >
          <CRB>Ждем ответа</CRB>
        </div>
      </Dialog>
    </>
  );
}

export default PageLoading;