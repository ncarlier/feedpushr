

export const config = {
  root: process.env.REACT_APP_API_ROOT || `${window.location.origin}/v1`
}

export const handleErrors = function(res) {
  if (!res.ok) {
    return res.json().then(json => Promise.reject(json))
  }
  return res.status === 200 ? res.json() : Promise.resolve()
}