import { User } from 'oidc-client'
import React, { createContext, ReactNode, useContext, useEffect, useState } from 'react'
import { addRequestInterceptor } from '../helpers/fetchAPI'
import { OIDCClient } from '../helpers/oidc-client'
import { ConfigContext } from './ConfigContext'
import oidcInterceptor from '../helpers/oidc-interceptor'

const AuthNContext = createContext<User | null>(null)

interface Props {
  children: ReactNode
}

const AuthNProvider = ({ children }: Props) => {
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<Error>()
  const [user, setUser] = useState<User | null>(null)
  const { _links } = useContext(ConfigContext)

  useEffect(() => {
    const { origin, href, pathname } = document.location
    const redirect = encodeURIComponent(href)
    const doAuth = async () => {
      const { issuer } = _links
      if (!issuer) {
        // no delegated authentication
        setLoading(false)
        return
      }
      const client = new OIDCClient({
        authority: issuer.href,
        client_id: 'feedpushr-ui',
        redirect_uri: `${origin}${pathname}signin-callback.html?redirect=${redirect}`,
        silent_redirect_uri: `${origin}${pathname}silent-renew.html`,
        post_logout_redirect_uri: origin + pathname,
      })
      addRequestInterceptor(oidcInterceptor(client))
      try {
        let user = await client.getUser()
        if (user === null) {
          return await client.login()
        } else if (user.expired) {
          user = await client.renewToken()
        }
        setUser(user)
      } catch (e) {
        setError(e)
      } finally {
        setLoading(false)
      }
    }
    doAuth()
  }, [])

  return (
    <AuthNContext.Provider value={user}>
      {error && <p>ERROR: {error.message}</p>}
      {loading && <p>Connecting...</p>}
      {!error && !loading && children}
    </AuthNContext.Provider>
  )
}

export { AuthNContext, AuthNProvider }
