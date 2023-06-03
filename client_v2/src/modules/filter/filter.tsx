import React, {ChangeEvent, useState} from 'react';
import classes from "./filter.module.css";
import {useAnalysis} from "../../stores/analysis-store";
import CDescription from "../../ui/description/description";
import {useFilterStore} from "../../stores/filter-store";
import {Logger} from "../../common/logger";
import btn_class from "../../ui/btn.module.css";
import Description from "../../ui/description/description";
import {Drawer, Input, InputNumber, Slider, Space, Switch} from "antd";
import {Placement} from "../../stores/placement";

interface IFilter {

}

const Filter = () => {
  // stores
  const analysisStore = useAnalysis()
  const filterStore = useFilterStore()

  const [title, setTitle] = useState("")
  const [description, setDescription] = useState("")
  const [price, setPrice]
    = useState<[number, number]>([filterStore.minPrice, filterStore.maxPrice])

  return (
    <>
      {
        analysisStore.analysis.length != 0 ?
          <div
            className={classes.filter_block}
          >
            <button
              onClick={filterStore.open}
              className={btn_class.btn}
            >
              Открыть фильтр
            </button>
            <CDescription>
              Фильтруйте результаты поиска по категориям
            </CDescription>
          </div> : <></>
      }

      <Drawer

        title="Фильтр"
        placement={Placement.PLACEMENT_BOTTOM}
        onClose={filterStore.close}
        open={filterStore.isOpen}
      >
        <div
          className="flex flex-col 2xl:flex-row xl:flex-row lg:flex-row scroll-auto"
        >
          <Input
            onChange={(event: ChangeEvent<HTMLInputElement>) => {
              setTitle(event.target.value)
            }}
            placeholder="Искать в названии"
            size="large"
          />
          <Description>
            Поле ввода для поиска в названии анализа.
          </Description>

          <Input
            onChange={(event: ChangeEvent<HTMLInputElement>) => {
              setDescription(event.target.value)
            }}
            placeholder="Искать в описание"
            size="large"
          />
          <Description>
            Поле ввода для поиска в описании анализа.
          </Description>
          <Space
            style={{ width: '100%' }}
          >
            <label>От</label>
            <InputNumber
              value={price[0]}
              defaultValue={price[0]}
              onChange={(value: number | null) => price[0] = value!!}
            />
            <label>До</label>
            <InputNumber
              value={price[1]}
              defaultValue={price[1]}
              onChange={(value: number | null) => price[1] = value!!}
            />
          </Space>
          <Description>
            Выбор диапазоны цены для фильтрации, чтобы не фильтровать по цене -
            оставьте максимально и минимальное значение 0.
          </Description>

          <button
            className={btn_class.btn}
            onClick={() => {
              const filterData = {
                title: title,
                description: description,
                price: price
              }
              Logger.Info("filter/to_store", filterData)
              filterStore.setQuery(filterData)

              filterStore.close()
            }}
          >
            Применить фильтр
          </button>
        </div>
      </Drawer>
    </>
  );
};

export default Filter;