import React from 'react';
import {IAnalysis} from "../../models/analysis";
import classes from "./style.module.css";
import {truncate} from "../../common/truncate";
import Description from "../../ui/description/description";

interface ICardAnalysis {
  openSelectCallback: (value: IAnalysis) => void,
  analysis: IAnalysis,
  colorLab: string
}

function CardAnalysis(props: ICardAnalysis) {
  return (
    <div
      onClick={() => props.openSelectCallback(props.analysis)}
      className={classes.card}
    >
      <h5
        className={classes.title}
      >
        {truncate(props.analysis.name, 50)}
      </h5>
      <p className={classes.description}>
        {truncate(props.analysis.description)}
      </p>
      <Description
        className="text-center mt-2"
      >
        Нажми, чтобы узнать подробнее
      </Description>
    </div>
  );
}

export default CardAnalysis;