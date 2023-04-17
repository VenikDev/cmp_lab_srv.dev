export interface ICoords {
  lat: string
  lon: string
}

export interface ICity {
  coords: ICoords,
  district: string
  name: string
  population: number
  subject: string
}