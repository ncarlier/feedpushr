
export interface Spec {
  name: string
  desc: string
  props: PropsSpec[]
}

export interface PropsSpec {
  name: string
  desc: string
  type: string
  options?: Object
}

export interface Props {
  [key: string]: any
}

interface Base {
  id: string
  alias: string
  name: string
  desc: string
  props: Props
  condition: string
  enabled: boolean
  nbSuccess: number
  nbError: number
}

interface BaseForm {
  alias: string
  name: string
  enabled: boolean
  condition: string
  props: Props
}

export type Output = Base & {
  filters?: Filter[]
}

export type OutputForm = BaseForm & {
  id?: string
}

export type Filter = Base

export type FilterForm = BaseForm & {
  output: string
  idx?: number
}

export type Entity = Base & {
  parentId?: string
  type: 'output' | 'filter'
}
