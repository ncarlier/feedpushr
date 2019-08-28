import React from 'react'

import Loader from '../common/Loader'
import Message from '../common/Message'
import matchResponse from '../helpers/matchResponse'
import { useAPI, usePageTitle } from '../hooks'
import OutputList from './OutputList'
import { Output } from './Types'

type OutputsResponse = Output[]

export default () => {
  usePageTitle('outputs')
  const [loading, outputs, error] = useAPI<OutputsResponse>('/outputs')

  const render = matchResponse<OutputsResponse>({
    Loading: () => (<Loader />),
    Data: data => (<OutputList outputs={data} />),
    Error: err => (<Message message={`Unable to retrieve outputs: ${err.message}`} variant="error" />)
  })

  return (<>{render(loading, outputs, error)}</>)
}
