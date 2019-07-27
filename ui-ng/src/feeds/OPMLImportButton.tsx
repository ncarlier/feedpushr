import React, { useContext } from 'react'

import { Tooltip } from '@material-ui/core'
import {
  CloudUpload as CloudUploadIcon
} from '@material-ui/icons'

import UploadButton from '../common/UploadButton'
import fetchAPI from '../helpers/fetchAPI'
import { MessageContext } from '../context/MessageContext'
import Message from '../common/Message'
import { withRouter, RouteComponentProps } from 'react-router'

interface Props {
  style?: React.CSSProperties
}

export default withRouter(({style, history} : Props & RouteComponentProps) => {
  const { showMessage } = useContext(MessageContext)
  
  const handleOnSelectFile = async (file: File) => {
    const formData = new FormData()
    formData.append('file', file)

    try {
      const res = await fetchAPI('/opml', null, {method: 'POST', body: formData})
      if (res.ok) {
        showMessage(<Message variant="success"  message={'OPML file imported'} />)
        history.push('/')
      } else {
        const err = await res.json()
        throw new Error(err.detail || res.statusText)
      }
    } catch (err) {
      showMessage(<Message variant="error"  message={`Unable to import OPML file: ${err.message}`} />)
    }
  }

  return (
    <Tooltip title="Import from OPML file" style={style}>
      <UploadButton size="small" variant="contained" color="default" onSelectFile={handleOnSelectFile}>
        Import
        <CloudUploadIcon />
      </UploadButton>
    </Tooltip>
  )
})
