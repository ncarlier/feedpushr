import React from 'react'
import { RouteComponentProps } from 'react-router'

import Loader from '../common/Loader'
import Message from '../common/Message'
import matchResponse from '../helpers/matchResponse'
import { useAPI, usePageTitle } from '../hooks'
import FeedList from './FeedList'
import { FeedPage } from './Types'

type FeedsResponse = FeedPage

type Props = RouteComponentProps

const buildReqFromQuery = (query: string) => {
  const params = new URLSearchParams(query)
  let page = 1
  if (params.has('page')) {
    page = parseInt(params.get('page') || '1', 10)
    if (isNaN(page)) {
      page = 1
    }
  }
  return { page, limit: 100 }
}

export default ({ location }: Props) => {
  const req = buildReqFromQuery(location.search)
  usePageTitle('feeds')
  const [loading, data, error] = useAPI<FeedsResponse>('/feeds', req)

  const render = matchResponse<FeedsResponse>({
    Loading: () => <Loader />,
    Data: (page) => <FeedList page={page} />,
    Error: (err) => <Message message={`Unable to retrieve feeds: ${err.message}`} variant="error" />,
  })

  return <>{render(loading, data, error)}</>
}
