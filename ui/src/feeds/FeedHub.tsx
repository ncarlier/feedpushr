import { Chip, Tooltip } from '@mui/material'
import { Cloud as CloudIcon } from '@mui/icons-material'

import { Feed } from './Types'

interface Props {
  feed: Feed
}

export default ({ feed }: Props) => {
  if (feed.hubUrl) {
    return (
      <Tooltip title="WebSub ready" sx={{ margin: 1 }}>
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
