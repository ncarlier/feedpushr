import React from 'react'
import { format } from 'timeago.js'

import { Tooltip } from '@material-ui/core'

interface Props {
  value?: string
  title?: React.ReactNode;
}

export default ({value, title}: Props) => {
  if (value) {
    return (
      <Tooltip title={title || value}>
        <span>{format(value!)}</span>
      </Tooltip>
    )
  }
  return (<span>-</span>)
}
