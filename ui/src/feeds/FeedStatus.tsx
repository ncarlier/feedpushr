import React from 'react'

import { makeStyles, Theme, Tooltip } from '@material-ui/core'
import { green } from '@material-ui/core/colors'

import classNames from '../helpers/classNames'
import { Feed } from './Types'

const useStyles = makeStyles((theme: Theme) => ({
  status: {
    display: 'inline-block',
    alignItems: 'center',
    backgroundColor: theme.palette.primary.main,
    color: theme.palette.primary.contrastText,
    minWidth: '2em',
    minHeight: '2em',
    padding: '.5em!important',
    lineHeight: '1em',
    borderRadius: '500rem',
    textAlign: 'center',
  },
  error: {
    backgroundColor: theme.palette.error.dark,
  },
  success: {
    backgroundColor: green[600],
  },
}))

interface Props {
  feed: Feed
}

export default ({ feed }: Props) => {
  const classes = useStyles()
  let $status = <div className={classNames(classes.status)}>0</div>
  let title = 'No feed aggregated'
  if (feed.errorCount) {
    title = feed.errorMsg ? feed.errorMsg : 'unexpected error'
    $status = <div className={classNames(classes.status, classes.error)}>{feed.errorCount}</div>
  } else if (feed.nbProcessedItems) {
    title = 'Aggregation success'
    $status = <div className={classNames(classes.status, classes.success)}>{feed.nbProcessedItems}</div>
  }
  return <Tooltip title={title}>{$status}</Tooltip>
}
