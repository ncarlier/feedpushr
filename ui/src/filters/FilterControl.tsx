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

  async function switchFilterStatus(event: React.ChangeEvent, check: boolean) {
    const update = {...filter, enabled: check, tags: filter.tags ? filter.tags.join(',') : '' }
    try {
      const res = await fetchAPI(`/filters/${filter.id}`, null, {
        method: 'PUT',
        body: JSON.stringify(update),
      })
      if (res.ok) {
        setStatus(check)
        showMessage(<Message variant="success" message={`Filter ${filter.name} ${check ? 'enabled' : 'disabled'}`} />)
      } else {
        throw new Error(res.statusText)
      }
    }
    catch (err) {
      showMessage(<Message variant="error"  message={`Unable to update filter: ${err.message}`} />)
    }
  }

  return (
    <Tooltip title="Enable/Disable">
      <Switch
        color="primary"
        checked={status}
        value={filter.enabled}
        onChange={switchFilterStatus}
      />
    </Tooltip>
  )
}
