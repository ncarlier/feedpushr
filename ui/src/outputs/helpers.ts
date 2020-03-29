import { Entity, Filter, Output } from './Types'

const desc = (name: string, alias: string, type: string) => `${alias ? alias : name} ${type}`

export const descEntity = (entity: Entity) => {
  const { name, alias, type } = entity
  return desc(name, alias, type)
}

export const descOutput = (output: Output) => {
  const { name, alias } = output
  return desc(name, alias, 'output')
}

export const descFilter = (filter: Filter) => {
  const { name, alias } = filter
  return desc(name, alias, 'filter')
}
