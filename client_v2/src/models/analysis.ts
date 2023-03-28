export interface IAnalysis {
  name: string
  description: string
  price: number
}

export interface IListAnalysis {
  analysis: IAnalysis[],
  isLoading: boolean,
  addAnalysis: (analysis: IAnalysis[]) => void
  changeStateLoading: () => void
}

export interface IFavoriteAnalysis {
  list: IAnalysis[]
  length: number
  addToFavorite: (analysis: IAnalysis) => void
}