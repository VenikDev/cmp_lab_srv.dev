import React from 'react';
import Title from "../../modules/title/title";
import SearchBlock from "../../modules/search-block/search-block";
import Carousel from "../../modules/carousel/carousel";
import Popular from "../../modules/popular/popular";
import CDescription from "../../ui/description/description";
import CRB from "../../ui/text/bold-red";

function Home() {
  return (
    <>
      <Title/>
      {/* Для поиска */}
      <SearchBlock/>
      {/* Для часто искомого */}
      <Popular/>
      <Carousel/>
    </>
  );
}

export default Home;