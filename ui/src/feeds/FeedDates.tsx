import React from 'react'

import TimeAgo from '../common/TimeAgo'
import { Feed } from './Types'

interface Props {
  feed: Feed
}

const toUTCString = (value: string) => {
  const date = new Date(value)
  return date.toUTCString()
}

export default ({ feed }: Props) => (
  <TimeAgo
    title={
      <dl>
        <dt>created</dt>
        <dd>{toUTCString(feed.cdate)}</dd>
        {!!feed.mdate && (
          <>
            <dt>updated</dt>
            <dd>{toUTCString(feed.mdate)}</dd>
          </>
        )}
        {!!feed.lastCheck && (
          <>
            <dt>last check</dt>
            <dd>{toUTCString(feed.lastCheck)}</dd>
          </>
        )}
        {!!feed.nextCheck && (
          <>
            <dt>next check</dt>
            <dd>{toUTCString(feed.nextCheck)}</dd>
          </>
        )}
      </dl>
    }
    value={feed.nextCheck}
  />
)
