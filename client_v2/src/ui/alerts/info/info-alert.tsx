import React from 'react';

interface IInfoAlert {
  children: string
  className?: string
}

function CInfoAlert(props: IInfoAlert) {
  return (
    <div className={`items-center bg-blue-300 p-3 rounded-md ${props.className}`}>
      <span className="text-blue-700">
        { props.children }
      </span>
    </div>
  );
}

export default CInfoAlert;