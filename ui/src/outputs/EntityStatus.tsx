import React from 'react'

import { makeStyles, Theme } from '@material-ui/core'
import { green } from '@material-ui/core/colors'

import classNames from '../helpers/classNames'
import { Entity } from './Types'

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
  entity: Entity
  error?: boolean
}

export default ({ entity, error = false }: Props) => {
  const classes = useStyles()
  if (error) {
    if (entity.nbError > 0) {
      return <div className={classNames(classes.status, classes.error)}>{entity.nbError}</div>
    }
  } else if (entity.nbSuccess > 0) {
    return <div className={classNames(classes.status, classes.success)}>{entity.nbSuccess}</div>
  }
  return <span>-</span>
}
