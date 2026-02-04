import { Box } from '@mui/material'
import { green } from '@mui/material/colors'

import { Entity } from './Types'

interface Props {
  entity: Entity
  error?: boolean
}

export default ({ entity, error = false }: Props) => {
  if (error) {
    if (entity.nbError > 0) {
      return (
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
          {entity.nbError}
        </Box>
      )
    }
  } else if (entity.nbSuccess > 0) {
    return (
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
        {entity.nbSuccess}
      </Box>
    )
  }
  return <span>-</span>
}
