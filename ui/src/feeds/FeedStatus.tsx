import { Box, Tooltip } from '@mui/material'
import { green } from '@mui/material/colors'

import { Feed } from './Types'

interface Props {
  feed: Feed
}

export default ({ feed }: Props) => {
  let $status = (
    <Box
      component="div"
      sx={(theme) => ({
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
      })}
    >
      0
    </Box>
  )
  let title = 'No feed aggregated'
  if (feed.nbErrors) {
    title = feed.lastErrorMsg ? feed.lastErrorMsg : 'unexpected error'
    $status = (
      <Box
        component="div"
        sx={(theme) => ({
          display: 'inline-block',
          alignItems: 'center',
          backgroundColor: theme.palette.error.dark,
          color: theme.palette.primary.contrastText,
          minWidth: '2em',
          minHeight: '2em',
          padding: '.5em!important',
          lineHeight: '1em',
          borderRadius: '500rem',
          textAlign: 'center',
        })}
      >
        {feed.nbErrors}
      </Box>
    )
  } else if (feed.nbProcessedItems) {
    title = 'Aggregation success'
    $status = (
      <Box
        component="div"
        sx={{
          display: 'inline-block',
          alignItems: 'center',
          backgroundColor: green[600],
          color: '#fff',
          minWidth: '2em',
          minHeight: '2em',
          padding: '.5em!important',
          lineHeight: '1em',
          borderRadius: '500rem',
          textAlign: 'center',
        }}
      >
        {feed.nbProcessedItems}
      </Box>
    )
  }
  return <Tooltip title={title}>{$status}</Tooltip>
}
