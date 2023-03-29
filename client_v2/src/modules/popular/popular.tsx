import React, {useEffect, useState} from 'react';
import {IPopular} from "./model";
import ky from "ky";
import {HOST_V1} from "../../net/consts";
import Carousel, { slidesToShowPlugin } from '@brainhubeu/react-carousel';
import '@brainhubeu/react-carousel/lib/style.css';
import classes from "./popular.module.css"
import CDescription from "../../ui/description/description";
import CRB from "../../ui/text/bold-red";

function Popular() {
  const [popular, setPopular] = useState<IPopular[]>()

  useEffect(() => {
    const getPopular = async () => {
      await ky(HOST_V1+"/get_popular")
        .json<IPopular[]>()
        .then(value => setPopular(value))
    }

    getPopular()
  }, [])

  return (
    <div
      className="w-full my-4"
    >
      <Carousel
        plugins={[
          'centered',
          'infinite',
          {
            resolve: slidesToShowPlugin,
            options: {
              numberOfSlides: 4,
            }
          },
        ]}
      >
        {
          popular && popular.map((item, idx) =>
            <div
              className={classes.item_carousel}
              key={idx}
            >
              {/* name */}
              <div
                className="text-center"
              >
                <CRB>{ item.name }</CRB>
              </div>
              {/* count */}
              <div
                className="text-center"
              >
                Искали <CRB>{ item.count }</CRB> раз
              </div>

              <CDescription
                className="text-xs text-center"
              >
                Нажните, чтобы искать
              </CDescription>
            </div>
          )
        }
      </Carousel>
    </div>
  );
}

export default Popular;