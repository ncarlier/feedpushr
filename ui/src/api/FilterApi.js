import {config, handleErrors} from './common'

export class FilterApi {
  constructor() {
    this.root = `${config.root}/filters`
  }

  list() {
    return fetch(this.root, {
      method: 'GET',
    }).then(handleErrors)
  }
}

const instance = new FilterApi()
export default instance