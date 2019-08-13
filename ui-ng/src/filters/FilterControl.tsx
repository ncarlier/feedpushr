import React, { useContext, useState } from 'react'

import { Switch, Tooltip } from '@material-ui/core'

import Message from '../common/Message'
import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'
import { Filter } from './Types'

interface Props {
  filter: Filter
}

export default ({filter}: Props) => {
  const [status, setStatus] = useState(filter.enabled)
  const { showMessage } = useContext(MessageContext)

  const switchFilterStatus = (event: React.ChangeEvent, check: boolean) => {
    const update = {enabled: check, ...filter}
    fetchAPI(`/filters/${filter.id}`, null, {
      method: 'PUT',
      body: JSON.stringify(update),
    })
    .then(res => {
      setStatus(check)
      showMessage(<Message variant="success"  message={`Filter ${filter.name} ${check ? 'enabled' : 'disabled'}`} />)
    }).catch(console.error)
  }

  return (
    <Tooltip title="Enabled/Disabled filter">
      <Switch
        color="primary"
        checked={status}
        value={filter.enabled}
        onChange={switchFilterStatus}
      />
    </Tooltip>
  )
}
