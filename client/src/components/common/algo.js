/**
 * Кол-во ключей в объекте
 * @param object {Object} - объект
 * @return {number} - Кол-во ключей
 */
const getCountKeys = (object) => {
  return Object.keys(object).length;
}

/**
 * Есть ли содержимое в объекте
 * @param object {Object} - объект для проверки
 * @return {boolean} - true, если пустой
 */
export function objectIsEmpty(object) {
  return getCountKeys(object) === 0
}