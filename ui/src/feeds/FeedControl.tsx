import React, { useContext, useEffect, useState } from 'react'

import { Switch } from '@material-ui/core'

import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'
import { Feed } from './Types'

interface Props {
  feed: Feed
}

export default ({ feed }: Props) => {
  const [status, setStatus] = useState(false)
  const { showMessage } = useContext(MessageContext)

  useEffect(() => {
    setStatus(feed.status === 'running')
  }, [feed])

  const switchFeedStatus = (event: React.ChangeEvent, check: boolean) => {
    const action = check ? 'start' : 'stop'
    fetchAPI(`/feeds/${feed.id}/${action}`, null, { method: 'POST' })
      .then((/*res*/) => {
        setStatus(check)
        showMessage(`${feed.title} feed is ${check ? 'running' : 'stopped'}`)
      })
      .catch(console.error)
  }

  return <Switch color="primary" title="start/stop" checked={status} value={feed.status} onChange={switchFeedStatus} />
}
