import React, {ChangeEvent} from 'react';
import CInput from "../../ui/input/input";
import classes from "./filter.module.css";
import {useAnalysis} from "../../stores/analysis-store";
import CDescription from "../../ui/description/description";
import Select from "../../ui/select/select";
import ExpendedCard from "../../ui/expended-card/expended-card";
import {useFilterStore} from "../../stores/filter-store";
import {Logger} from "../../common/logger";

interface IFilter {

}

const Filter = () => {
  // stores
  const analysisStore = useAnalysis()
  const filterStore = useFilterStore()

  const list = [
    "Искать в описании",
    "Искать в заголовке"
  ]

  const callbackSelect = (value: string) => {
    Logger.Info("callbackSelect", value)
    filterStore.setCategory(value)
  }

  const callbackInput = (value: string) => {
    Logger.Info("callbackInput", value)
    filterStore.setQuery(value)
  }

  return (
    analysisStore.analysis.length != 0 ?
    <div
      className={classes.filter_block}
    >
      <ExpendedCard
        title="Фильтрация результатов"
      >
        <div
          className="flex"
        >
          <Select
            id="filter_to"
            list={list}
            callbackSelect={callbackSelect}
          />
          <CInput
            onInput={(event: ChangeEvent<HTMLInputElement>) => callbackSelect(event.target.value) }
            placeholder="Фильтр"
          />
        </div>
        <CDescription>
          Фильтруйте результаты поиска по категориям
        </CDescription>
      </ExpendedCard>
    </div> : ""
  );
};

export default Filter;