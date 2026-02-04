import { IconButton, Tooltip } from '@mui/material'
import { Public as PublicIcon } from '@mui/icons-material'

import { Feed } from './Types'

interface Props {
  feed: Feed
}

export default ({ feed }: Props) => {
  if (feed.htmlUrl) {
    return (
      <Tooltip title="Open website in a new tab" sx={{ marginRight: 1 }}>
        <IconButton
          aria-label="public"
          sx={{ marginRight: 1 }}
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
