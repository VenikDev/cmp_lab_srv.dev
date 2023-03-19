import React from 'react';

interface IAlertError {
  children: string
}

function CAlertError(props: IAlertError) {
  return (
    <div className="items-center bg-red-300 p-3 rounded-md">
      <span className="text-red-700">
        { props.children }
      </span>
    </div>
  );
}

export default CAlertError;