const API_ROOT = process.env.REACT_APP_API_ROOT || window.location.origin

export default async (uri: string, params: any = {}, init: RequestInit) => {
  const url = new URL(`${API_ROOT}/v1${uri}`)
  if (params) {
    Object.keys(params).forEach(key => url.searchParams.append(key, params[key]))
  }
  return await fetch(url.toString(), init)
}
