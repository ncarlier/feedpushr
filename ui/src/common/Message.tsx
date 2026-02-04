import { SnackbarContent, IconButton, Box } from '@mui/material'

import {
  CheckCircle as CheckCircleIcon,
  Warning as WarningIcon,
  Error as ErrorIcon,
  Info as InfoIcon,
  Close as CloseIcon,
} from '@mui/icons-material'

import { amber, green } from '@mui/material/colors'

const variantIcon = {
  success: CheckCircleIcon,
  warning: WarningIcon,
  error: ErrorIcon,
  info: InfoIcon,
}

export interface Props {
  className?: string
  text: string
  onClose?: () => void
  variant: keyof typeof variantIcon
}

export default (props: Props) => {
  const { className, text, onClose, variant, ...other } = props
  const Icon = variantIcon[variant]

  return (
    <SnackbarContent
      sx={(theme) => ({
        backgroundColor:
          variant === 'success'
            ? green[600]
            : variant === 'warning'
            ? amber[700]
            : variant === 'info'
            ? theme.palette.primary.main
            : variant === 'error'
            ? theme.palette.error.dark
            : undefined,
      })}
      className={className}
      aria-describedby="client-snackbar"
      message={
        <Box
          component="span"
          id="client-snackbar"
          sx={{
            display: 'flex',
            alignItems: 'center',
          }}
        >
          <Icon sx={{ fontSize: 20, opacity: 0.9, marginRight: 1 }} />
          {text}
        </Box>
      }
      action={[
        <IconButton key="close" aria-label="close" color="inherit" onClick={onClose}>
          <CloseIcon sx={{ fontSize: 20 }} />
        </IconButton>,
      ]}
      {...other}
    />
  )
}
