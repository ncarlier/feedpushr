import React, { createContext, ReactNode, useEffect, useState } from 'react'
import fetchAPI from '../helpers/fetchAPI'

interface LinkType {
  href: string
}

interface ConfigContextType {
  version: string
  _links: Record<string, LinkType>
}

const defaultConfig: ConfigContextType = {
  version: 'snapshot',
  _links: {},
}

const ConfigContext = createContext<ConfigContextType>(defaultConfig)

interface Props {
  children: ReactNode
}

const ConfigProvider = ({ children }: Props) => {
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<Error>()
  const [config, setConfig] = useState<ConfigContextType>(defaultConfig)

  useEffect(() => {
    const abortController = new AbortController()
    const doFetchAPI = async () => {
      try {
        const res = await fetchAPI('/', null, { signal: abortController.signal })
        if (res.ok) {
          const json = await res.json()
          setConfig(json)
          return
        }
        throw new Error(res.statusText)
      } catch (e) {
        setError(e)
      } finally {
        setLoading(false)
      }
    }
    doFetchAPI()
    return () => abortController.abort()
  }, [])

  return (
    <ConfigContext.Provider value={config}>
      {error && <p>ERROR: {error.message}</p>}
      {loading && <p>LOADING...</p>}
      {!error && !loading && children}
    </ConfigContext.Provider>
  )
}

export { ConfigContext, ConfigProvider }
