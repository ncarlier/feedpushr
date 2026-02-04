import { useContext } from 'react'

import { Tooltip, Button } from '@mui/material'
import { CloudDownload as CloudDownloadIcon } from '@mui/icons-material'

import fetchAPI from '../helpers/fetchAPI'
import { MessageContext } from '../context/MessageContext'

export default () => {
  const { showMessage } = useContext(MessageContext)

  const handleOnClick = async () => {
    try {
      const res = await fetchAPI('/opml', null, { method: 'GET' })
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
      showMessage(`Unable to export feeds to OPML file: ${(err as Error).message}`, 'error')
    }
  }

  return (
    <Tooltip title="Export to OPML file">
      <Button size="small" variant="contained" onClick={handleOnClick} endIcon={<CloudDownloadIcon />}>
        Export
      </Button>
    </Tooltip>
  )
}
