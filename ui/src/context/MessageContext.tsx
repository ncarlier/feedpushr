import React, { createContext, ReactNode, useState } from 'react'

type MessageVariant = 'success' | 'warning' | 'error' | 'info'

interface Message {
  text: string
  variant: MessageVariant
}

interface MessageContextType {
  message: Message
  showMessage: (text: string, variant?: MessageVariant) => void
}

const MessageContext = createContext<MessageContextType>({
  message: { text: '', variant: 'info' },
  showMessage: () => true,
})

interface Props {
  children: ReactNode
}

const MessageProvider = ({ children }: Props) => {
  const [message, setMessage] = useState<Message>({ text: '', variant: 'info' })

  const showMessage = (text: string, variant: MessageVariant = 'success') => {
    setMessage({ text, variant })
  }

  return <MessageContext.Provider value={{ message, showMessage }}>{children}</MessageContext.Provider>
}

export { MessageContext, MessageProvider }
