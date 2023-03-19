import React, {useState} from 'react';
import {CButton} from "../../ui/button/button";
import classes from './search-block.module.css'
import Dialog from "../../ui/dialog/dialog";

function SearchBlock() {
  const [visibleDialog, setVisibleDialog] = useState(false)

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
        <h1 className="block font-bold">
          Title
        </h1>
        <hr className="my-2"/>
        <span className="block">
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Accusantium,
          aut blanditiis consectetur consequatur doloremque dolores, eius fugiat
          fugit harum illum magni maiores nostrum obcaecati porro provident
          suscipit tenetur. Sit, vitae.
        </span>
        <button
          onClick={() => setVisibleDialog(false)}
          className="mx-auto mt-2 bg-red-500 text-white p-2 rounded-md"
        >
          Закрыть
        </button>
      </Dialog>
    </>
  );
}

export default SearchBlock;