import React from 'react'

import { useAPI } from '../hooks'
import matchResponse from '../helpers/matchResponse'
import FeedList from './FeedList'
import { Feed } from './Types'
import Message from '../common/Message'
import Loader from '../common/Loader'

type FeedsResponse = Feed[]

export default () => {
  const [loading, feeds, error] = useAPI<FeedsResponse>('/feeds', {page: 1, limit: 100})

  const render = matchResponse<FeedsResponse>({
    Loading: () => (<Loader />),
    Data: data => (<FeedList feeds={data} />),
    Error: err => (<Message message={`Unable to retrieve feeds: ${err.message}`} variant="error" />)
  })

  return (<>{render(loading, feeds, error)}</>)
}
