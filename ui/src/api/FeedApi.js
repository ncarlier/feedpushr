import { config, handleErrors } from './common'

const defaultHeaders = new Headers()
defaultHeaders.append('Content-Type', 'application/json')

export class FeedApi {
  constructor() {
    this.root = `${config.root}/feeds`
  }

  list(page = 1, limit = 100) {
    const url = new URL(this.root)
    const params = { page, limit }
    Object.keys(params).forEach(key => url.searchParams.append(key, params[key]))
    return fetch(url, {
      method: 'GET',
      headers: defaultHeaders,
    }).then(handleErrors)
  }

  add(url, tags, title) {
    const _url = new URL(this.root)
    const params = { url, tags: tags.join(","), title }
    Object.keys(params).forEach(key => _url.searchParams.append(key, params[key]))
    return fetch(_url, {
      method: 'POST',
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      }
    }).then(handleErrors)
  }
  
  update(feed, payload) {
    const _url = new URL(`${this.root}/${feed.id}`)
    const { tags, title } = payload
    const params = { tags: tags.join(","), title }
    Object.keys(params).forEach(key => _url.searchParams.append(key, params[key]))
    return fetch(_url, {
      method: 'PUT',
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      }
    }).then(handleErrors)
  }

  remove(id) {
    return fetch(`${this.root}/${id}`, {
      method: 'DELETE',
    }).then(handleErrors)
  }

  start(id) {
    return fetch(`${this.root}/${id}/start`, {
      method: 'POST',
    }).then(handleErrors)
  }

  stop(id) {
    return fetch(`${this.root}/${id}/stop`, {
      method: 'POST',
    }).then(handleErrors)
  }

}

const instance = new FeedApi()
export default instance
