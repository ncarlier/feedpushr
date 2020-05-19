import React, { useContext } from 'react'
import { RouteComponentProps, withRouter } from 'react-router'

import {
  Tooltip,
  Dialog,
  useTheme,
  useMediaQuery,
  DialogContent,
  DialogActions,
  Button,
  DialogTitle,
} from '@material-ui/core'
import { CloudUpload as CloudUploadIcon } from '@material-ui/icons'

import Message from '../common/Message'
import UploadButton from '../common/UploadButton'
import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'
import OPMLImportJobStatus from './OPMLImportJobStatus'

interface ImportJobResult {
  id: string
}

interface Props {
  style?: React.CSSProperties
}

export default withRouter(({ style, history }: Props & RouteComponentProps) => {
  const [open, setOpen] = React.useState(false)
  const [jobID, setJobID] = React.useState('')
  const theme = useTheme()
  const fullScreen = useMediaQuery(theme.breakpoints.down('sm'))
  const { showMessage } = useContext(MessageContext)

  const handleClose = () => {
    setOpen(false)
    history.push('/')
    history.replace('/feeds')
  }

  const handleOnSelectFile = async (file: File) => {
    const formData = new FormData()
    formData.append('file', file)

    try {
      const res = await fetchAPI('/opml', null, { method: 'POST', body: formData })
      if (!res.ok) {
        const err = await res.json()
        throw new Error(err.detail || res.statusText)
      }
      const { id } = (await res.json()) as ImportJobResult
      setJobID(id)
      setOpen(true)
    } catch (err) {
      showMessage(<Message variant="error" message={`Unable to import OPML file: ${err.message}`} />)
    }
  }

  return (
    <>
      <Tooltip title="Import from OPML file" style={style}>
        <UploadButton size="small" variant="contained" color="default" onSelectFile={handleOnSelectFile}>
          Import
          <CloudUploadIcon />
        </UploadButton>
      </Tooltip>
      <Dialog fullScreen={fullScreen} open={open} onClose={handleClose} aria-labelledby="import-dialog-title">
        <DialogTitle id="import-dialog-title">{'Import OPML file'}</DialogTitle>
        <DialogContent>
          <OPMLImportJobStatus jobID={jobID} />
        </DialogContent>
        <DialogActions>
          <Button autoFocus onClick={handleClose} color="primary">
            Dismiss
          </Button>
        </DialogActions>
      </Dialog>
    </>
  )
})
