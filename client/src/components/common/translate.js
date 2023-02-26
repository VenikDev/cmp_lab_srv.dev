import * as keys from "../../assets/dict/key_word.json";

/**
 *
 * @param searchKey {string} - ключ
 * @constructor
 */
export function Translate(searchKey) {
  console.log("key = ", searchKey, ", value ", keys[searchKey])
  return keys[searchKey]
}