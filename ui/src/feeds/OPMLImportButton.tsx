import { useContext, useState } from 'react'
import { useNavigate } from 'react-router-dom'

import {
  Tooltip,
  Dialog,
  useTheme,
  useMediaQuery,
  DialogContent,
  DialogActions,
  Button,
  DialogTitle,
} from '@mui/material'
import { CloudUpload as CloudUploadIcon } from '@mui/icons-material'

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

export default ({ style }: Props) => {
  const navigate = useNavigate()
  const [open, setOpen] = useState(false)
  const [jobID, setJobID] = useState('')
  const theme = useTheme()
  const fullScreen = useMediaQuery(theme.breakpoints.down('sm'))
  const { showMessage } = useContext(MessageContext)

  const handleClose = () => {
    setOpen(false)
    navigate('/feeds', { replace: true })
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
      showMessage(`Unable to import OPML file: ${(err as Error).message}`, 'error')
    }
  }

  return (
    <>
      <Tooltip title="Import from OPML file" style={style}>
        <UploadButton size="small" variant="contained" onSelectFile={handleOnSelectFile} endIcon={<CloudUploadIcon />}>
          Import
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
}
