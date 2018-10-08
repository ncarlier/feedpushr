import { config, handleErrors } from './common'

export class OutputApi {
  constructor() {
    this.root = `${config.root}/output`
  }

  get() {
    return fetch(this.root, {
      method: 'GET',
    }).then(handleErrors)
  }
}

const instance = new OutputApi()
export default instance