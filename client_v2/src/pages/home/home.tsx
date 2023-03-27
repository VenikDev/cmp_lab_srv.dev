import React from 'react';
import Title from "../../modules/title/title";
import SearchBlock from "../../modules/search-block/search-block";
import Carousel from "../../modules/carousel/carousel";

function Home() {
  return (
    <>
      <Title/>
      <SearchBlock/>
      <Carousel/>
    </>
  );
}

export default Home;