// export type LabAndAnalysis = Map<string, IAnalysis[]>

import {v4 as uuid4} from "uuid";

interface NameAndListAnalysis {
  nameLab: string
  list: IAnalysis[]
}

export interface LabAndAnalysis {
  name_lab: string
  list: IAnalysis[]
}

export interface IAnalysis {
  id: string
  name: string
  description: string
  price: number
  original_url: string
  isSelect: boolean
}

export interface IListAnalysis {
  analysis: LabAndAnalysis[],
  isLoading: boolean,
  addAnalysis: (newAnalysis: LabAndAnalysis[]) => void
  changeStateLoading: () => void
}