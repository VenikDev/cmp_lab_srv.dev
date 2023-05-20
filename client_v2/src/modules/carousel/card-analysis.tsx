import React from 'react';
import {IAnalysis} from "../../models/analysis";
import classes from "./style.module.css";
import {truncate} from "../../common/truncate";
import Description from "../../ui/description/description";
import {useFilterStore} from "../../stores/filter-store";
import {FiltrationTypes} from "../../ui/expended-card/FiltrationTypes";

interface ICardAnalysis {
  openSelectCallback: (value: IAnalysis) => void,
  analysis: IAnalysis,
  colorLab: string
}

export default CardAnalysis;

function CardAnalysis(props: ICardAnalysis) {
  // stores
  const filterStore = useFilterStore()


  const isValidTest = (lookForPlace: string, query: string) => {
    const pattern = new RegExp(`${query}`);
    return pattern.test(lookForPlace);
  }

  return (
    (filterStore.category == FiltrationTypes.SEARCH_DESCRIPTION &&
      isValidTest(props.analysis.description, filterStore.query) ||
      filterStore.category == FiltrationTypes.SEARCH_TITLE &&
      isValidTest(props.analysis.name, filterStore.query)) ?
    <div
        onClick={() => props.openSelectCallback(props.analysis)}
        className={classes.card}
    >
        <h5
            className={classes.title}
        >
          { truncate(props.analysis.name, 50) }
        </h5>
        <p className={classes.description}>
          { truncate(props.analysis.description) }
        </p>
        <Description
            className="text-center mt-2"
        >
            Нажми, чтобы узнать подробнее
        </Description>
    </div> : <></>
  );
}
