import React, {useState} from 'react';
import Dialog from "../../ui/dialog/dialog";
import {useAnalysis} from "../../stores/analysis-store";
import waitingGif from "../../assets/waiting.gif"
import CRB from "../../ui/text/bold-red";

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
          src={waitingGif}
          alt="Не могу прогрузить веселую гифку"
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