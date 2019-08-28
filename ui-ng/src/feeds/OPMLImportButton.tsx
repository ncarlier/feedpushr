import React, { useContext } from 'react'
import { RouteComponentProps, withRouter } from 'react-router'

import { Tooltip } from '@material-ui/core'
import { CloudUpload as CloudUploadIcon } from '@material-ui/icons'

import Message from '../common/Message'
import UploadButton from '../common/UploadButton'
import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'

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
        history.replace('/feeds')
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
