import React from 'react'

import { Chip } from '@material-ui/core'
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles'

import { Feed } from './Types'

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    tag: {
      marginRight: theme.spacing(0.5),
    },
  }),
)

interface Props {
  feed: Feed
}

export default ({feed: {tags = []}}: Props) => {
  const classes = useStyles()
  return (
    <>{ tags.map(tag => <Chip key={tag} label={tag} className={classes.tag} />) }</>
  )
}
