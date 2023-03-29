import React from "react";
import classes from './App.module.css'
import NavBar from "../../modules/nav-bar/nav-bar";
import PageLoading from "../../modules/page-loaging/page-loading";
import {BrowserRouter, Route, Routes} from "react-router-dom";
// pages
import Home from "../home/home";
import Footer from "../../modules/footer/footer";
import About from "../about/about";
import Favorite from "../favorite/favorite";

function App() {

  return (
    <div className={classes.App}>
      <BrowserRouter>
        <NavBar/>
        <Routes>
          {/* main */}
          <Route
            path="/"
            element={ <Home/> }
          />

          {/* about */}
          <Route
            path="/about"
            element={ <About/> }
          />

          {/* favorite */}
          <Route
            path="/favorite"
            element={ <Favorite/> }
          />
        </Routes>
        <Footer/>
        <PageLoading/>
      </BrowserRouter>
    </div>
  )
}

export default App
