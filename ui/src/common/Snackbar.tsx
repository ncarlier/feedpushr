import React, { useContext, SyntheticEvent, useEffect, useState } from 'react'

import { Snackbar } from '@material-ui/core'

import { MessageContext } from '../context/MessageContext'
import Message from './Message'

export default () => {
  const { message } = useContext(MessageContext)
  const [open, setOpen] = useState(false)

  useEffect(() => {
    setOpen(message.text !== '')
  }, [message])

  const handleClose = (event?: SyntheticEvent, reason?: string) => {
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
      <Message text={message.text} variant={message.variant} onClose={handleClose} />
    </Snackbar>
  )
}
