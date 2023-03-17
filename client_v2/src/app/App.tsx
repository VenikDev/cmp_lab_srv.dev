import React, { useRef } from "react";
import classes from './App.module.css'
import {CButton} from '../component/ui/button/button'
import {CInput} from "../component/ui/input/input";

function App() {

  return (
    <div className={classes.App}>
      <h2>
        React icon
      </h2>
      <CButton>
        Нажми на меня
      </CButton>
      <CInput
        placeholder={"Text"}
      />
    </div>
  )
}

export default App
