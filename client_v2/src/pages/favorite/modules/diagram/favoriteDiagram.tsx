import React, {FunctionComponent, PureComponent} from 'react';
import {
  Area,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip, ResponsiveContainer, BarChart, Legend, Bar
} from "recharts";
import {curveCardinal} from "d3-shape";
import {useFavorite} from "../../../../stores/favorit-store";
import CDescription from "../../../../ui/description/description";

const cardinal = curveCardinal.tension(0.2);

const CustomTooltip: FunctionComponent = () => {
  return (
    <>

    </>
  )
}

const FavoriteDiagram: FunctionComponent = () => {
  const favoriteStore = useFavorite()

  return (
    <>
      <ResponsiveContainer width='100%' aspect={4.0 / 3.0}>
        <BarChart
          data={favoriteStore.selectedList}
          margin={{
            top: 10,
            right: 30,
            left: 0,
            bottom: 0
          }}
        >
          <CartesianGrid strokeDasharray="3 3"/>
          <XAxis dataKey="analysis.name"/>
          <YAxis/>
          <Tooltip/>
          <Tooltip content={<CustomTooltip/>}/>
          <Legend/>
          <Bar
            type="monotone"
            dataKey="analysis.price"
            name="цена (в руб.)"
            stroke="#8884d8"
            fill="#8884d8"
            fillOpacity={0.3}
          />
        </BarChart>
      </ResponsiveContainer>
      <CDescription>
        Нажмите на колонку, чтобы посмотреть название анализа и его цену
      </CDescription>
    </>
  );
}

export default FavoriteDiagram