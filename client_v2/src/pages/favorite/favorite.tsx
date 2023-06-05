import React from 'react';
import {useFavorite} from "../../stores/favorit-store";
import FavoriteDiagram from "./modules/diagram/favoriteDiagram";

function Favorite() {
  // get store
  const favoriteStore = useFavorite()

  return (
    <div>
      { favoriteStore.selectedList.map(item =>
        <div
          className="my-5 border-2 border-main-border"
        >
          <h1>
            { item.name }
          </h1>
          <span>
            { item.price }
          </span>\
          <FavoriteDiagram
            title="Избранное"
          />
        </div>
      ) }
    </div>
  );
}

export default Favorite;