import React from 'react'

import { makeStyles, Theme } from '@material-ui/core'
import { green } from '@material-ui/core/colors'

import classNames from '../helpers/classNames'
import { Filter } from './Types'

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
  }
}))

interface Props {
  filter: Filter
  error?: boolean
}

export default ({filter, error = false}: Props) => {
  const classes = useStyles()
  if (error) {
    if (filter.props.nbError && filter.props.nbError > 0) {
      return <div className={classNames(classes.status, classes.error)}>{filter.props.nbError}</div>
    }
  } else if (filter.props.nbSuccess && filter.props.nbSuccess > 0) {
    return <div className={classNames(classes.status, classes.success)}>{filter.props.nbSuccess}</div>
  }
  return <span>-</span>
}
