export type LabAndAnalysis = Map<string, IAnalysis[]>

export interface IAnalysis {
  name: string
  description: string
  price: number
  original_url: string
}

export interface IListAnalysis {
  analysis: LabAndAnalysis | null,
  isLoading: boolean,
  addAnalysis: (newAnalysis: LabAndAnalysis) => void
  changeStateLoading: () => void
}

export interface IFavoriteAnalysis {
  list: IAnalysis[]
  length: number
  addToFavorite: (analysis: IAnalysis) => void
}