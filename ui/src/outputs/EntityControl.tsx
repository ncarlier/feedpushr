import React, { useContext, useState } from 'react'

import { Switch, Tooltip } from '@material-ui/core'

import Message from '../common/Message'
import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'
import { Entity } from './Types'
import { descEntity } from './helpers'

interface Props {
  entity: Entity
}

export default ({ entity }: Props) => {
  const [status, setStatus] = useState(entity.enabled)
  const { showMessage } = useContext(MessageContext)

  let url = `/outputs/${entity.id}`
  if (entity.type === 'filter') {
    url = `/outputs/${entity.parentId}/filters/${entity.id}`
  }

  async function switchEntityStatus(event: React.ChangeEvent, check: boolean) {
    const update = { ...entity, enabled: check }
    try {
      const res = await fetchAPI(url, null, {
        method: 'PUT',
        body: JSON.stringify(update),
      })
      if (res.ok) {
        setStatus(check)
        showMessage(<Message variant="success" message={`${descEntity(entity)} ${check ? 'enabled' : 'disabled'}`} />)
      } else {
        throw new Error(res.statusText)
      }
    } catch (err) {
      showMessage(<Message variant="error" message={`Unable to update ${descEntity(entity)}: ${err.message}`} />)
    }
  }

  return (
    <Tooltip title="Enable/Disable">
      <Switch color="primary" checked={status} value={entity.enabled} onChange={switchEntityStatus} />
    </Tooltip>
  )
}
