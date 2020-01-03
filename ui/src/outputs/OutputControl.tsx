import React, { useContext, useState } from 'react'

import { Switch, Tooltip } from '@material-ui/core'

import Message from '../common/Message'
import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'
import { Output } from './Types'

interface Props {
  output: Output
}

export default ({output}: Props) => {
  const [status, setStatus] = useState(output.enabled)
  const { showMessage } = useContext(MessageContext)

  async function switchOutputStatus(event: React.ChangeEvent, check: boolean) {
    const update = {...output, enabled: check}
    try {
      const res = await fetchAPI(`/outputs/${output.id}`, null, {
        method: 'PUT',
        body: JSON.stringify(update),
      })
      if (res.ok) {
        setStatus(check)
        showMessage(<Message variant="success" message={`Output ${output.name} ${check ? 'enabled' : 'disabled'}`} />)
      } else {
        throw new Error(res.statusText)
      }
    }
    catch (err) {
      showMessage(<Message variant="error"  message={`Unable to update output: ${err.message}`} />)
    }
  }

  return (
    <Tooltip title="Enable/Disable">
      <Switch
        color="primary"
        checked={status}
        value={output.enabled}
        onChange={switchOutputStatus}
      />
    </Tooltip>
  )
}
