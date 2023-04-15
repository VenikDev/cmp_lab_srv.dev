import React from 'react';
import Title from "../../modules/title/title";
import SearchBlock from "../../modules/search-block/search-block";
import CCarousel from "../../modules/carousel/CCarousel";
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
      <CDescription
        className="text-center"
      >
        Потяните <CRB>влево</CRB> или <CRB>вправо</CRB>, чтоб посмотреть еще
      </CDescription>
      <CCarousel/>
    </>
  );
}

export default Home;