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
    filterStore.setQuery(value)
  }

  function handlerInput(event: ChangeEvent<HTMLInputElement>) {
    Logger.Info("handlerInput", event.target.value)

    callbackInput(event.target.value)
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
          className="flex flex-col 2xl:flex-row xl:flex-row lg:flex-row"
        >
          <Select
            id="filter_to"
            list={list}
            callbackSelect={callbackSelect}
          />
          <CInput
            onInput={handlerInput}
            placeholder="Фильтр"
          />
        </div>
        <CDescription>
          Фильтруйте результаты поиска по категориям
        </CDescription>
      </ExpendedCard>
    </div> : <></>
  );
};

export default Filter;