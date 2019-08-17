import React from 'react'

import Loader from '../common/Loader'
import Message from '../common/Message'
import matchResponse from '../helpers/matchResponse'
import { useAPI, usePageTitle } from '../hooks'
import FeedList from './FeedList'
import { Feed } from './Types'

type FeedsResponse = Feed[]

export default () => {
  usePageTitle('feeds')
  const [loading, feeds, error] = useAPI<FeedsResponse>('/feeds', {page: 1, limit: 100})

  const render = matchResponse<FeedsResponse>({
    Loading: () => <Loader />,
    Data: data => <FeedList feeds={data} />,
    Error: err => <Message message={`Unable to retrieve feeds: ${err.message}`} variant="error" />
  })

  return (<>{render(loading, feeds, error)}</>)
}
