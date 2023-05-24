// export type LabAndAnalysis = Map<string, IAnalysis[]>

interface NameAndListAnalysis {
  nameLab: string
  list: IAnalysis[]
}

export interface LabAndAnalysis {
  name_lab: string
  list: IAnalysis[]
}

export interface IAnalysis {
  name: string
  description: string
  price: number
  original_url: string
}

export interface IListAnalysis {
  analysis: LabAndAnalysis[],
  isLoading: boolean,
  addAnalysis: (newAnalysis: LabAndAnalysis[]) => void
  changeStateLoading: () => void
}