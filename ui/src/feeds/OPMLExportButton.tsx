import React, { useContext } from 'react'

import { Tooltip, Button } from '@material-ui/core'
import {
  CloudDownload as CloudDownloadIcon
} from '@material-ui/icons'

import fetchAPI from '../helpers/fetchAPI'
import { MessageContext } from '../context/MessageContext'
import Message from '../common/Message'

export default () => {
  const { showMessage } = useContext(MessageContext)
  
  const handleOnClick = async () => {
    try {
      const res = await fetchAPI('/opml', null, {method: 'GET'})
      if (res.ok) {
        const body = await res.text()
        const element = document.createElement('a')
        element.setAttribute('href', 'data:text/xml;charset=utf-8,' + encodeURIComponent(body))
        element.setAttribute('download', 'my-feedpushr-feeds.opml')
        element.style.display = 'none'
        document.body.appendChild(element)
        element.click()
        document.body.removeChild(element)
      } else {
        const err = await res.json()
        throw new Error(err.detail || res.statusText)
      }
    } catch (err) {
      showMessage(<Message variant="error"  message={`Unable to export feeds to OPML file: ${err.message}`} />)
    }
  }

  return (
    <Tooltip title="Export to OPML file">
      <Button size="small" variant="contained" color="default" onClick={handleOnClick}>
        Export
        <CloudDownloadIcon />
      </Button>
    </Tooltip>
  )
}
