import { OIDCClient } from './oidc-client'

const interceptor = (client: OIDCClient) => async (init: RequestInit) => {
  let user = await client.getUser()
  if (user === null) {
    throw new Error('user not logged in')
  }
  if (user.expired) {
    user = await client.renewToken()
  }
  if (user.access_token) {
    const headers = new Headers(init.headers)
    headers.set('Authorization', `Bearer ${user.access_token}`)
    init.headers = headers
  }
  return init
}

export default interceptor
