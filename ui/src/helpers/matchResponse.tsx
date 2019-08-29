import React, { ReactNode } from 'react'

export interface ResponsePattern<T> {
  Loading: () => ReactNode
  Error: (err: Error) => ReactNode
  Data: (data: T) => ReactNode
}

function matchResponse<T>(
  p: ResponsePattern<T>
): (loading: boolean, data?: T, error?: Error) => ReactNode {
  return (loading: boolean, data?: T, error?: Error): ReactNode => {
    return (
      <>
        {loading && p.Loading()}
        {error && p.Error(error)}
        {data && p.Data(data)}
      </>
    )
  }
}

export default matchResponse
