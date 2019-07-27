import React from 'react'

import { Chip } from '@material-ui/core'

interface Props {
  value: string[]
}

export default ({value = []}: Props) => (
  <>{ value.map(tag => <Chip key={tag} label={tag} />) }</>
)
