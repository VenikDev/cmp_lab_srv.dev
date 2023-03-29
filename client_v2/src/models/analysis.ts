export type LabAndAnalysis = Map<string, IAnalysis[]>

export interface IAnalysis {
  name: string
  description: string
  price: number
  original_url: string
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