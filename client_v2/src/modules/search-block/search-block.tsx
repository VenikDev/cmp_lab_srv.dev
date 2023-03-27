import React, {useEffect, useState} from 'react';
import CDescription from "../../ui/description/description";
import Dialog from "../../ui/dialog/dialog";
import CInput from "../../ui/input/input";
import {AiOutlineSearch, TfiClose} from "react-icons/all";
import ky from "ky";
import {HOST_V1, TypeRequest} from "../../common/net";
import CCheckBox from "../../ui/radio-btn/radio-btn";
import CAlertError from "../../ui/alerts/error/alert-error";
import CRB from "../../ui/text/bold-red";
import {Key} from "../../common/keys";
import {useAnalysis} from "../../stores/analysis-store";

function SearchBlock() {
  // Для открытия/закрытия диалогового окна
  const [visibleDialog, setVisibleDialog] = useState(false)
  const [labs, setLabs] = useState<string[]>()
  const analysisStore = useAnalysis()

  // get names of labs
  useEffect(() => {
    const getLabs = async () => {
      const options = {
        method: TypeRequest.GET,
      };
      await ky(`${HOST_V1}get_names_labs`, options)
        .json<string[]>().then(value => {
          setLabs(value)
        });
    }
    getLabs()
  }, [])

  // when sending a request
  const sendReq = async () => {
    // hide dialog with input field
    setVisibleDialog(false)

    // show loader
    analysisStore.changeStateLoading()
  }

  // listen "Esc" button
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
          placeholder="Поиск анализа"
        />
        {/* Описание */}
        <CDescription>
          Введите <CRB>ключевое слово</CRB>, по которому нужно найти интерсующий анализ
        </CDescription>
        {/* Выбор лабораторий */}
        <div className="ml-2">
          {
            labs && labs.length != 0 ?
              labs?.map((lab, idx) =>
                <>
                  <CCheckBox
                    key={idx}
                    value={lab}
                    id={`${lab}-${idx}`}
                    name="labs"
                    label={lab}
                  />
                </>
              )
              : <CAlertError>
                Ошибка: не возможно найти список лабораторий
              </CAlertError>
          }
        </div>
        <CDescription>
          Выберите <CRB>лаборатории</CRB>, которые вас интересуют
        </CDescription>
        {/* Отправка запроса для поиска по клоючевому слову */}
        <button
          onClick={() => sendReq()}
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