import React from "react";
import classes from './App.module.css'
import NavBar from "../../modules/nav-bar/nav-bar";
import Title from "../../modules/title/title";
import SearchBlock from "../../modules/search-block/search-block";
import Carousel from "../../modules/carousel/carousel";

function App() {

  return (
    <div className={classes.App}>
      <NavBar/>
      <Title/>
      <SearchBlock/>
      <Carousel/>
    </div>
  )
}

export default App
