import React from 'react';
import Autosuggest from 'react-autosuggest';
import {IPopular} from "../popular/model";

const SearchField = () => {
  function renderItem(item: IPopular) {
    return (
      <div
        className="flex"
      >
        <div>
          { item.name }
        </div>
        <div>
          { item.count }
        </div>
      </div>
    )
  }

  return (
    <>
      <Autosuggest
        suggestions={suggestions}
        onSuggestionsFetchRequested={this.onSuggestionsFetchRequested}
        onSuggestionsClearRequested={this.onSuggestionsClearRequested}
        getSuggestionValue={getSuggestionValue}
        renderSuggestion={renderItem}
        inputProps={inputProps}
      />
    </>
  );
};

export default SearchField;