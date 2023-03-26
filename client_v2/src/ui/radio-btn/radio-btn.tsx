import React from 'react';

interface ICheckPoint {
  value: string
  id: string
  name: string
  label: string
}

function CCheckBox(props: ICheckPoint) {
  return (
    <div className="block my-1">
      <input
        className=""
        type="checkbox"
        id={props.id}
        name={props.name}
        value={props.value}
        // checked={true}
      />
      <label
        className="ml-2"
        htmlFor={props.id}
      >
        {props.label}
      </label>
    </div>
  );
}

export default CCheckBox;