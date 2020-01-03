import React, { ReactNode } from 'react'

interface Props {
  href: string
  children: ReactNode
}

export default ({href, children}: Props) => (
  <a href={`https://github.com/ncarlier/feedpushr/blob/master/${href}`} target="_blank" rel="noopener noreferrer">
    {children}
  </a>
)
