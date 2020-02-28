import React from 'react'

import { Chip, Tooltip, Theme } from '@material-ui/core'
import { Cloud as CloudIcon } from '@material-ui/icons'
import { makeStyles, createStyles } from '@material-ui/styles'

import { Feed } from './Types'

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    margin: {
      margin: theme.spacing(1),
    },
  }),
);

interface Props {
  feed: Feed
}

export default ({ feed }: Props) => {
  const classes = useStyles()
  if (!!feed.hubUrl) {
    return (
      <Tooltip title="PubSubHubbud ready" className={classes.margin}>
        <Chip
          variant="outlined"
          size="small"
          label="hub"
          component="a"
          href={feed.hubUrl}
          clickable
          icon={<CloudIcon />}
        />
      </Tooltip>
    )
  }
  return null
}
