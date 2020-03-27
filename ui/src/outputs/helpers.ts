import { Entity } from "./Types"

export const descEntity = (entity: Entity) => {
  const name = entity.alias ? entity.alias : entity.name
  return `${name} ${entity.type}`
}
