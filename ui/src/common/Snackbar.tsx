import { useContext, useEffect, useState } from 'react'

import { Snackbar, SnackbarCloseReason } from '@mui/material'

import { MessageContext } from '../context/MessageContext'
import Message from './Message'

export default () => {
  const { message } = useContext(MessageContext)
  const [open, setOpen] = useState(false)

  useEffect(() => {
    setOpen(message.text !== '')
  }, [message])

  const handleClose = (_event: any, reason?: SnackbarCloseReason) => {
    if (reason === 'clickaway') {
      return
    }
    setOpen(false)
  }

  return (
    <Snackbar
      anchorOrigin={{
        vertical: 'bottom',
        horizontal: 'left',
      }}
      open={open}
      autoHideDuration={5000}
      onClose={handleClose}
    >
      <div>
        <Message text={message.text} variant={message.variant} onClose={() => setOpen(false)} />
      </div>
    </Snackbar>
  )
}
