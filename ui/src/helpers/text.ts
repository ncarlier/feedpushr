export const headline = (value: string) => value.slice(0, value.indexOf('.') + 1)

export const afterHeadline = (value: string) => value.slice(value.indexOf('.') + 1)
