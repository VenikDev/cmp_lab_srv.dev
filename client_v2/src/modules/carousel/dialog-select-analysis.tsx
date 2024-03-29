import React from 'react';
import CRB from "../../ui/text/strong_bold";
import classes from "./style.module.css";
import {useSelectAnalysis} from "../../stores/select-analysis-store";
import Dialog from "../../ui/dialog/dialog";

function DialogSelectAnalysis() {
  const selectAnalysisStore = useSelectAnalysis()

  return (
    <Dialog
      open={selectAnalysisStore.isOpen}
      title="Подробности анализа"
      callbackClose={selectAnalysisStore.changeState}
    >
      <h1>
        <CRB>
          {selectAnalysisStore.analysis?.name}
        </CRB>
      </h1>
      <hr
        className="my-3"
      />
      <h3>
        {selectAnalysisStore.analysis?.description}
      </h3>
      <div
        className="text-right mt-2"
      >
        <h1
          className={classes.price_analysis}
        >
          {selectAnalysisStore.analysis?.price} руб.
        </h1>
      </div>
      <a
        className={classes.link_to_lab}
        href={selectAnalysisStore.analysis?.original_url}
        target="_blank"
      >
        Перейти на сайт лаборатории
      </a>
      <br/>
      {/*<MapWrapper/>*/}
    </Dialog>
  );
}

export default DialogSelectAnalysis;