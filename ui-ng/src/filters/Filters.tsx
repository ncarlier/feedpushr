import React from 'react'

import Loader from '../common/Loader'
import Message from '../common/Message'
import matchResponse from '../helpers/matchResponse'
import { useAPI } from '../hooks'
import FilterList from './FilterList'
import { Filter } from './Types'

type FeedsResponse = Filter[]

export default () => {
  const [loading, filters, error] = useAPI<FeedsResponse>('/filters')

  const render = matchResponse<FeedsResponse>({
    Loading: () => (<Loader />),
    Data: data => (<FilterList filters={data} />),
    Error: err => (<Message message={`Unable to retrieve filters: ${err.message}`} variant="error" />)
  })

  return (<>{render(loading, filters, error)}</>)
}
