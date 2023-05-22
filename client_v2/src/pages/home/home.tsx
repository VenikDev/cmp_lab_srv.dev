import React from 'react';
import Title from "../../modules/title/title";
import SearchBlock from "../../modules/search-block/search-block";
import Carousel from "../../modules/carousel/carousel";
import Popular from "../../modules/popular/popular";
import Filter from "../../modules/filter/filter";
import Footer from "../../modules/footer/footer";

function Home() {
  return (
    <>
      <Title/>
      {/* Для поиска */}
      <SearchBlock/>
      {/* Для часто искомого */}
      <Popular/>
      <Filter/>
      <Carousel/>
      <Footer/>
    </>
  );
}

export default Home;