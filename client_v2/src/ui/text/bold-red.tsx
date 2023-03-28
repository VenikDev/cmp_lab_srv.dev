import React from 'react';

interface IBoldRed {
  children: React.ReactNode
}

function CBoldRed(props: IBoldRed) {
  return (
    <b className="text-red-700">
      { props.children }
    </b>
  );
}

export default CBoldRed;