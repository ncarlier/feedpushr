import React from 'react'

import { CircularProgress, makeStyles } from '@material-ui/core'

const useStyles = makeStyles((/*theme: Theme*/) => ({
  loader: {
    display: 'flex',
    padding: '20px',
    justifyContent: 'center',
  },
}))

export default () => {
  const classes = useStyles()
  return (
    <div className={classes.loader}>
      <CircularProgress disableShrink />
    </div>
  )
}
