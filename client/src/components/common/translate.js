import * as keys from '../../assets/dict/key_word.json';

// This code defines a function called "Translate", which takes a single argument
// called "searchKey". It returns the value associated with "searchKey" in the "keys"
// object. The "keys" object is assumed to be defined elsewhere in the code.
// The purpose of the function appears to be to translate a given search term into
// its corresponding value in a lookup table.
export function Translate(searchKey) {
  return keys[searchKey];
}
