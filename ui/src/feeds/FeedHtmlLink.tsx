import React from 'react'

import { IconButton, Theme, Tooltip } from '@material-ui/core'
import { Public as PublicIcon } from '@material-ui/icons'
import { makeStyles, createStyles } from '@material-ui/styles'

import { Feed } from './Types'

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    margin: {
      marginRight: theme.spacing(1),
    },
  }),
);

interface Props {
  feed: Feed
}

export default ({ feed }: Props) => {
  const classes = useStyles()
  if (!!feed.htmlUrl) {
    return (
      <Tooltip title="Open website in a new tab" className={classes.margin}>
        <IconButton
          aria-label="public"
          className={classes.margin}
          size="small"
          component="a"
          href={feed.htmlUrl}
          target="_blank"
          >
            <PublicIcon fontSize="inherit" />
        </IconButton>
      </Tooltip>      
    )
  }
  return null
}
