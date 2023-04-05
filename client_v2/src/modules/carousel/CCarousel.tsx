import React, {useEffect} from 'react';
import {useAnalysis} from '../../stores/analysis-store'
import Carousel, {slidesToShowPlugin} from "@brainhubeu/react-carousel";
import classes from "../popular/popular.module.css";
import {getAnalysis} from "../../net/requests";
import CRB from "../../ui/text/bold-red";
import CDescription from "../../ui/description/description";


function CCarousel() {
  const analysisStore = useAnalysis()
  const result: React.ReactNode[] = []

  if (analysisStore.analysis)
  {
    for (const [key, value] of analysisStore.analysis) {
      result.push(
        value.map((item, key) =>
          <div key={key}>
            <h2>
              <a href={item.original_url}>
                { item.name }
              </a>
            </h2>
          </div>
        )
      )
    }
  }


  return (
    <div>
      { result }
    </div>
  );
}

export default CCarousel;