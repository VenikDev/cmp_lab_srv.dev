import React, {useEffect, useRef, useState} from 'react';
import CDescription from "../../ui/description/description";
import Dialog from "../../ui/dialog/dialog";
import CInput from "../../ui/input/input";
import {AiOutlineSearch, TfiClose} from "react-icons/all";
import CCheckBox from "../../ui/radio-btn/radio-btn";
import CAlertError from "../../ui/alerts/error/alert-error";
import CRB from "../../ui/text/bold-red";
import {Key} from "../../common/keys";
import ky from "ky";
import {HOST_V1} from "../../net/consts";
import {getAnalysis} from "../../net/requests";
import {useAnalysis} from "../../stores/analysis-store";
import {LabAndAnalysis} from "../../models/analysis";
import {useGlobalProperties} from "../../stores/global-properties-store";

function SearchBlock() {
  // Для открытия/закрытия диалогового окна
  const [visibleDialog, setVisibleDialog] = useState(false)
  // Для поисковой тсроки
  const [nameAna, setNameAna] = useState<string>()
  // название лабораторий
  const [labs, setLabs] = useState<string[]>()

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
    analysisStore.changeStateLoading()

    // let result = new Map<string, IAnalysis[]>()
    const analysis = await getAnalysis<LabAndAnalysis>(nameAna!!, globalPropertiesStore.selectCity)
    analysisStore.addAnalysis(analysis)
    console.log(analysis)
    analysisStore.changeStateLoading()
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
        className="w-full bg-red-500 text-white rounded-md p-2"
        onClick={() => setVisibleDialog(true)}
      >
        Открыть поиск
      </button>
      <Dialog
        open={visibleDialog}
        callbackClose={setVisibleDialog}
        title="Поиск"
      >
        {/* Для воода ключевого слова */}
        <CInput
          onInput={(event) => setNameAna(event.target.value) }
          placeholder="Поиск анализа"
        />
         Отправка запроса для поиска по клоючевому слову
        <button
          onClick={async () => {
            setVisibleDialog(false)
            await sendReq()

          }}
          className="mx-auto mt-2 bg-red-500 text-white p-2 rounded-md block px-5 flex"
        >
          <AiOutlineSearch className="w-5 h-5 mr-2 my-auto"/>
          Найти
        </button>
      </Dialog>
    </>
  );
}

export default SearchBlock;