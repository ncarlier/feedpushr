import { config } from './common'

function handleErrors(response) {
  if (!response.ok) {
    return response.json().then(json => Promise.reject(json))
  }
  return response.text()
}

export class OPMLApi {
  constructor() {
    this.root = `${config.root}/opml`
  }

  export() {
    return fetch(this.root, {
      method: 'GET',
    }).then(handleErrors)
  }

  import(file) {
    const formData = new FormData()
    formData.append('file', file)

    return fetch(this.root, {
      method: 'POST',
      body: formData
    }).then(handleErrors)
  }

}

const instance = new OPMLApi()
export default instance