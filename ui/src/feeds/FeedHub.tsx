import React from 'react'

import { Chip, Tooltip } from '@material-ui/core'
import { Cloud as CloudIcon } from '@material-ui/icons'

import { Feed } from './Types'

interface Props {
  feed: Feed
}

export default ({ feed }: Props) => {
  if (!!feed.hubUrl) {
    return (
      <Tooltip title="PubSubHubbud ready">
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
