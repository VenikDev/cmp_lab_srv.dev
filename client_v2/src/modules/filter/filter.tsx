import React, {useState} from 'react';
import CInput from "../../ui/input/input";
import classes from "./filter.module.css";
import {useAnalysis} from "../../stores/analysis-store";

interface IFilter {

}

const Filter = () => {
  // stores
  const analysisStore = useAnalysis()

  const [filterAnalysis, setFilterAnalysis] = useState<string>()


  return (
    analysisStore.analysis.length != 0 ?
    <div
      className={classes.filter_block}
    >
      <h1
        className="text-xl"
      >
        Фильтрация результатов
      </h1>
      <CInput
        onInput={(event) => setFilterAnalysis(event.target.value) }
        placeholder="Фильтр"
      />
    </div> : ""
  );
};

export default Filter;