export const API_ROOT =
  process.env.REACT_APP_API_ROOT || window.location.origin + window.location.pathname.replace(/\/ui(\/)*$/, '')

export type RequestInterceptor = (init: RequestInit) => Promise<RequestInit>

let _interceptors: RequestInterceptor[] = []

export const addRequestInterceptor = (...interceptors: RequestInterceptor[]): void => {
  _interceptors = [..._interceptors, ...interceptors]
}

export default async (uri: string, params: any = {}, init: RequestInit) => {
  for (const interceptor of _interceptors) {
    init = await interceptor(init)
  }
  const url = new URL(`${API_ROOT}/v2${uri}`)
  if (params) {
    Object.keys(params).forEach((key) => url.searchParams.append(key, params[key]))
  }
  return await fetch(url.toString(), init)
}
