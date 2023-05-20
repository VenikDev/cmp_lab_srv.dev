import React, {useEffect, useRef, useState} from 'react';
import Dialog from "../../ui/dialog/dialog";
import CInput from "../../ui/input/input";
import {AiOutlineSearch} from "react-icons/all";
import {Key} from "../../common/keys";
import ky from "ky";
import {HOST_V1} from "../../net/consts";
import {getAnalysis} from "../../net/requests";
import {useAnalysis} from "../../stores/analysis-store";
import {LabAndAnalysis} from "../../models/analysis";
import {useGlobalProperties} from "../../stores/global-properties-store";
import AlertError from "../../ui/alerts/error/alert-error";
import Description from "../../ui/description/description";
import classes from "./search-block.module.css"
import StrongBold from "../../ui/text/strong_bold";
import CDescription from "../../ui/description/description";

function SearchBlock() {
  // Для открытия/закрытия диалогового окна
  const [visibleDialog, setVisibleDialog] = useState(false)
  // Для поисковой тсроки
  const [nameAnalysis, setNameAnalysis] = useState<string>()
  // название лабораторий
  const [labs, setLabs] = useState<string[]>()

  const [error, setError] = useState<string>()
  // stores
  const globalPropertiesStore = useGlobalProperties()
  const analysisStore = useAnalysis()

  // get names of labs
  useEffect(() => {
    (async () => {
      await ky(HOST_V1+`/get_names_labs`)
        .json<string[]>().then(value => {
          setLabs(value)
        });
    })()
  }, [])

  const sendReq = async () => {
    if (!nameAnalysis || nameAnalysis.length == 0) {
      setError("Поле для запроса пустое")
      return false
    }
    else {
      setError(undefined)
    }

    analysisStore.changeStateLoading()
    const analysis = await getAnalysis<LabAndAnalysis[]>(nameAnalysis, globalPropertiesStore.selectCity?.name!!)
    analysisStore.addAnalysis(analysis)
    analysisStore.changeStateLoading()

    return true
  }

  // This is a function named "keyTest" that takes an event as input.
  // It uses a switch statement to check the value of the "key" property of the event.
  // If the key is "ESCAPE", it sets the variable "setVisibleDialog" to false.
  // If the key is "ENTER", it sets "setVisibleDialog" to false and waits for the "sendReq" function to finish before continuing.
  // The "Key" in "case Key.ENTER" is likely an enumeration with predefined values for different types of keys.
  const keyTest = async (event: any) => {
    switch (event.key) {
      case Key.ESCAPE: {
        setVisibleDialog(false)
        break
      }
      case Key.ENTER: {
        setVisibleDialog(false)
        await sendReq()
        break
      }
    }
  };

  useEffect(() => {
    document.addEventListener("keydown", keyTest, false);

    return () => {
      document.removeEventListener("keydown", keyTest, false);
    };
  }, [keyTest]);

  // create component
  return (
    <>
      <button
        className={classes.search_btn}
        onClick={() => setVisibleDialog(true)}
      >
        Открыть поиск
      </button>
      <CDescription>
        Наймите, чтоб открыть окно для <StrongBold>ввода запроса</StrongBold>
      </CDescription>
      <Dialog
        open={visibleDialog}
        callbackClose={setVisibleDialog}
        title="Поиск"
      >
        {/* Для воода ключевого слова */}
        <CInput
          onInput={(event) => setNameAnalysis(event.target.value) }
          placeholder="Поиск анализа"
        />
        {/* alert for error */}
        {
          error && <AlertError className="my-2">{error}</AlertError>
        }
        {/* description */}
        <Description>
          Отправка запроса для поиска по <StrongBold> ключевому слову </StrongBold>
        </Description>
        {/* send request */}
        <button
          onClick={async () => {
            const isSuccess = await sendReq()
            if (isSuccess) {
              setVisibleDialog(false)
            }
          }}
          className={classes.find_tests}
        >
          <AiOutlineSearch className="w-5 h-5 mr-2 my-auto"/>
          Найти
        </button>
      </Dialog>
    </>
  );
}

export default SearchBlock;