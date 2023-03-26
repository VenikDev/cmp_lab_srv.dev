export interface IAnalysis {

}

export interface IListAnalysis {
  analysis: IAnalysis[],
  isLoading: boolean,
  addAnalysis: (analysis: IAnalysis[]) => void
  changeStateLoading: () => void
}

export interface IFavoriteAnalysis {

}