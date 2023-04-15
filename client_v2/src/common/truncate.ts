export function truncate(str: string) {
  const length: number = 100
  return str.length > length ? str.substring(0, length) + "..." : str;
}