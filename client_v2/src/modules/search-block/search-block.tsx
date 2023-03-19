import React, {useEffect, useState} from 'react';
import CDescription from "../../ui/description/description";
import Dialog from "../../ui/dialog/dialog";
import CInput from "../../ui/input/input";
import {TfiClose} from "react-icons/all";
import ky from "ky";
import {HOST_V1, TypeRequest} from "../config/variables";
import CCheckPoint from "../../ui/radio-btn/radio-btn";
import CAlertError from "../../ui/alerts/error/alert-error";

function SearchBlock() {
  // Для открытия/закрытия диалогового окна
  const [visibleDialog, setVisibleDialog] = useState(false)
  const [labs, setLabs] = useState<string[]>()

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

  const sendReq = async () => {
    setVisibleDialog(false)
  }

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
      >
        <div className="flex justify-around">
          <h1 className="block font-bold grow">
            Поиск
          </h1>
          <button
            className="cursor-pointer"
            onClick={() => setVisibleDialog(false)}
          >
            <TfiClose/>
          </button>
        </div>
        <hr className="my-2"/>
        {/* Для воода ключевого слова */}
        <CInput
          placeholder="Поиск анализа"
        />
        {/* Описание */}
        <CDescription>
          Введите <b>ключевое слово</b>, по которому нужно найти интерсующий анализ
        </CDescription>
        {/* Выбор лабораторий */}
        <div className="ml-2">
          {
            labs != undefined && labs.length != 0 ?
              labs?.map((lab, idx) =>
                <>
                  <CCheckPoint
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
          Выберите на <b>лаборатории</b>, которые вас интересуют
        </CDescription>
        {/* Отправка запроса для поиска по клоючевому слову */}
        <button
          onClick={() => sendReq()}
          className="mx-auto mt-2 bg-red-500 text-white p-2 rounded-md block px-5"
        >
          Найти
        </button>
      </Dialog>
    </>
  );
}

export default SearchBlock;