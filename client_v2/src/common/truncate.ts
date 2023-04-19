export function truncate(str: string, limitation: number = 100) {
  return str.length > limitation ? str.substring(0, limitation) + "..." : str;
}