export function findMaxWithPred<T>(arr: T[], pred: (n: T) => boolean): T | undefined {
  if (arr.length === 0) {
    return undefined
  }
  let max: T = arr[0];

  for (let i = 0; i < arr.length; i++) {
    if (pred(arr[i])) {
      if (arr[i] > max) {
        max = arr[i];
      }
    }
  }

  return max;
}

export function findMinWithPred<T>(arr: T[], pred: (n: T) => boolean): T | undefined {
  if (arr.length === 0) {
    return undefined
  }
  let min: T = arr[0];

  for (let i = 0; i < arr.length; i++) {
    if (pred(arr[i])) {
      if (arr[i] < min) {
        min = arr[i];
      }
    }
  }

  return min;
}