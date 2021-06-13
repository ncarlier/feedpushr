import React, { useContext } from 'react'

import { Typography } from '@material-ui/core'

import { ConfigContext } from '../context/ConfigContext'

export default () => {
  const { version } = useContext(ConfigContext)
  return (
    <Typography align="right" color="textSecondary" variant="caption">
      {version}
    </Typography>
  )
}
