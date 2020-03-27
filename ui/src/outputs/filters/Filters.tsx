import React from 'react'

import Loader from '../../common/Loader'
import Message from '../../common/Message'
import matchResponse from '../../helpers/matchResponse'
import { useAPI, usePageTitle } from '../../hooks'
import FilterList from './FilterList'
import { Filter } from './Types'

type FiltersResponse = Filter[]

export default () => {
  usePageTitle('filters')
  const [loading, filters, error] = useAPI<FiltersResponse>('/filters')

  const render = matchResponse<FiltersResponse>({
    Loading: () => (<Loader />),
    Data: data => (<FilterList filters={data} />),
    Error: err => (<Message message={`Unable to retrieve filters: ${err.message}`} variant="error" />)
  })

  return (<>{render(loading, filters, error)}</>)
}
