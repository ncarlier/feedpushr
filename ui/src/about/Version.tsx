import { useContext } from 'react'

import { Typography } from '@mui/material'

import { ConfigContext } from '../context/ConfigContext'

export default () => {
  const { version } = useContext(ConfigContext)
  return (
    <Typography align="right" color="textSecondary" variant="caption">
      {version}
    </Typography>
  )
}
