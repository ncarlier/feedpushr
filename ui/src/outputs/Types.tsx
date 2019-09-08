
export interface OutputProps {
    nbSuccess?: number
    nbError?: number
    [key: string]: any
}

export interface Output {
  id: number
  alias: string
  name: string
  desc: string
  enabled: boolean
  tags: string[]
  props: OutputProps
}

export interface PropsSpec {
  name: string
  desc: string
  type: string
}

export interface OutputSpec {
  name: string
  desc: string
  props: PropsSpec[]
}

export interface OutputForm {
  id?: number
  alias: string
  name: string
  tags: string[]
  props: OutputProps
}
