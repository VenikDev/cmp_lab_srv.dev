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
import {useCityStore} from "../../stores/city-store";
import {useAnalysis} from "../../stores/analysis-store";
import {LabAndAnalysis} from "../../models/analysis";

function SearchBlock() {
  // Для открытия/закрытия диалогового окна
  const [visibleDialog, setVisibleDialog] = useState(false)
  // Для поисковой тсроки
  const [nameAna, setNameAna] = useState<string>()
  // название лабораторий
  const [labs, setLabs] = useState<string[]>()

  // stores
  const cityStore = useCityStore()
  const analysisStore = useAnalysis()

  // get names of labs
  useEffect(() => {
    const getLabs = async () => {
      await ky(HOST_V1+`/get_names_labs`)
        .json<string[]>().then(value => {
          setLabs(value)
        });
    }
    getLabs()
  }, [])

  const sendReq = async () => {
    analysisStore.changeStateLoading()

    // let result = new Map<string, IAnalysis[]>()
    const analysis = await getAnalysis<LabAndAnalysis>(nameAna!!, cityStore.city)
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
        {/* Описание */}
        {/*<CDescription>*/}
        {/*  Введите <CRB>ключевое слово</CRB>, по которому нужно найти интерсующий анализ*/}
        {/*</CDescription>*/}
        {/* Выбор лабораторий*/}
        {/*<div className="ml-2">*/}
        {/*  {*/}
        {/*    labs && labs.length != 0 ?*/}
        {/*      labs?.map((lab, idx) =>*/}
        {/*        <label className="relative inline-flex items-center cursor-pointer">*/}
        {/*          <input type="checkbox" value="" className="sr-only peer"/>*/}
        {/*            <div*/}
        {/*              className="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600"></div>*/}
        {/*            <span className="ml-3 text-sm font-medium text-gray-900 dark:text-gray-300">Toggle me</span>*/}
        {/*        </label>*/}
        {/*      )*/}
        {/*      : <CAlertError>*/}
        {/*        Ошибка: не возможно найти список лабораторий*/}
        {/*      </CAlertError>*/}
        {/*  }*/}
        {/*</div>*/}
        {/*<CDescription>*/}
        {/*  Выберите <CRB>лаборатории</CRB>, которые вас интересуют*/}
        {/*</CDescription>*/}
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