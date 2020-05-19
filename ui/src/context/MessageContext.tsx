import React, { createContext, ReactNode, SyntheticEvent, useState } from 'react'

import { Snackbar } from '@material-ui/core'

interface MessageContextType {
  showMessage: (message: ReactNode | null) => void
}

const MessageContext = createContext<MessageContextType>({
  showMessage: () => true,
})

interface Props {
  children: ReactNode
}

const MessageProvider = ({ children }: Props) => {
  const [message, setMessage] = useState<ReactNode | null>(null)
  const [open, setOpen] = useState(false)

  const handleClose = (event?: SyntheticEvent, reason?: string) => {
    if (reason === 'clickaway') {
      return
    }
    setOpen(false)
  }

  const showMessage = (msg: ReactNode | null) => {
    setMessage(msg)
    setOpen(true)
  }

  return (
    <MessageContext.Provider value={{ showMessage }}>
      {children}
      <Snackbar
        anchorOrigin={{
          vertical: 'bottom',
          horizontal: 'left',
        }}
        open={open}
        autoHideDuration={5000}
        onClose={handleClose}
      >
        <>{message}</>
      </Snackbar>
    </MessageContext.Provider>
  )
}

export { MessageContext, MessageProvider }
