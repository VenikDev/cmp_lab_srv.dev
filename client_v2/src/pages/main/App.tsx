import React, {useEffect} from "react";
import classes from './App.module.css'
import NavBar from "../../modules/nav-bar/nav-bar";
import PageLoading from "../../modules/page-loaging/page-loading";
import {BrowserRouter, Route, Routes} from "react-router-dom";
import {useGlobalProperties} from "../../stores/global-properties-store";
import {Logger} from "../../common/logger";

// pages
import Home from "../home/home";
import Footer from "../../modules/footer/footer";
import About from "../about/about";
import Favorite from "../favorite/favorite";


function App() {
  const globalPropertiesStore = useGlobalProperties()

  useEffect(() => {
    const isMobile = /iPhone|iPad|iPod|Android/i.test(navigator.userAgent);
    if (isMobile) {
      Logger.Info("isMobile", isMobile)
      globalPropertiesStore.setIsPhone(true)
    }
  }, [])

  return (
    <div className={classes.App}>
      <BrowserRouter>
        <main
          className="flex flex-col h-full"
        >
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
        </main>
        {/*<Footer/>*/}
        <PageLoading/>
      </BrowserRouter>
    </div>
  )
}

export default App
