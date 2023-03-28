import React from 'react';
import {useFavorite} from "../../stores/favorit-store";

function Favorite() {
  // get store
  const favoriteStore = useFavorite()

  return (
    <div>
      { favoriteStore.list.map(item =>
        <div>
          {  }
        </div>
      ) }
    </div>
  );
}

export default Favorite;