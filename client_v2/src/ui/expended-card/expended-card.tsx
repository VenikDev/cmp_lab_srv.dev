import React, {useState} from 'react';
import classes from "./expended-card.module.css"
import {IoIosArrowDown, IoIosArrowForward} from "react-icons/all";
import CDescription from "../description/description";
import StrongBold from "../text/strong_bold";

interface IExpendedCard {
  title: string
  children: React.ReactNode
}

const ExpendedCard = (props: IExpendedCard) => {
  const [isExpanded, setIsExpanded] = useState(false);

  const toggleExpanded = () => {
    setIsExpanded(!isExpanded);
  }

  return (
    <div
      className={`${classes.card} ${isExpanded ? classes.expanded : ""}`}
    >
      <div
        className="w-full flex items-center"
        onClick={() => toggleExpanded()}
      >
        {
          isExpanded ?
            <IoIosArrowDown
              className="w-5"
            />
            :
            <IoIosArrowForward
              className="w-5"
            />
        }
        <h1
          className={classes.title}
        >
          { props.title }
        </h1>
      </div>
      {
        isExpanded && (
          <div>
            { props.children }
          </div>
        )
      }
    </div>
  );
};

export default ExpendedCard;