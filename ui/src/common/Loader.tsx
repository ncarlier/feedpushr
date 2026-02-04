import { CircularProgress, Box } from '@mui/material'

export default () => {
  return (
    <Box
      sx={{
        display: 'flex',
        padding: '20px',
        justifyContent: 'center',
      }}
    >
      <CircularProgress disableShrink />
    </Box>
  )
}
