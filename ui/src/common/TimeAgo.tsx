import React from 'react'

import { Tooltip } from '@material-ui/core'
import { format } from 'timeago.js'

interface Props {
  value?: string
}

export default ({value}: Props) => {
  if (value) {
    return (
      <Tooltip title={value}>
        <span>{format(value!)}</span>
      </Tooltip>
    )
  }
  return (<span>-</span>)
}
