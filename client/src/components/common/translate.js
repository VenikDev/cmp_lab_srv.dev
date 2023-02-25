import * as keys from "../../assets/dict/key_word.json";


/**
 *
 * @param searchKey {string} - ключ
 * @constructor
 */
export function Translate(searchKey) {
  return keys[searchKey]
}