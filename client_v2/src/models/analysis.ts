export interface IAnalysis {

}

export interface IListAnalysis {
  analysis: IAnalysis[],
  isLoading: boolean,
  addAnalysis: (analysis: IAnalysis[]) => void
}

export interface IFavoriteAnalysis {

}