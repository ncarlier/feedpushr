import React, { useState, useContext } from 'react'

import { Switch, Tooltip } from '@material-ui/core'

import { Feed } from './Types'
import fetchAPI from '../helpers/fetchAPI'
import { MessageContext } from '../context/MessageContext'
import Message from '../common/Message'

interface Props {
  feed: Feed
}

export default ({feed}: Props) => {
  const [status, setStatus] = useState(feed.status === 'running')
  const { showMessage } = useContext(MessageContext)

  const switchFeedStatus = (event: React.ChangeEvent, check: boolean) => {
    const action = check ? 'start' : 'stop'
    fetchAPI(`/feeds/${feed.id}/${action}`, null, {method: 'POST'})
    .then(res => {
      setStatus(check)
      showMessage(<Message variant="success"  message={`Feed ${feed.title} ${check ? 'running' : 'stopped'}`} />)
    }).catch(console.error)
  }

  return (
    <Tooltip title="Start/Stop aggregator">
      <Switch
        color="primary"
        checked={status}
        value={feed.status}
        onChange={switchFeedStatus}
      />
    </Tooltip>
  )
}
