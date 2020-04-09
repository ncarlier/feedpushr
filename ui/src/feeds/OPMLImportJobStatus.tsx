import React from 'react'

import {
  List,
  ListItem,
  ListItemText,
  ListItemIcon
} from '@material-ui/core'

import {
  CheckCircle as SuccessIcon,
  Error as ErrorIcon
} from '@material-ui/icons'
import { green, red } from '@material-ui/core/colors'

import Message from '../common/Message'
import Loader from '../common/Loader'

const API_ROOT = process.env.REACT_APP_API_ROOT || window.location.origin

type EventSourceStatus = 'open' | 'closed' | 'error'

interface JobResultItem {
  XMLURL: string
  error?: string
}

interface Props {
  jobID: string
}

export default ({jobID}: Props) => {
  const [items, setItems] = React.useState<JobResultItem[]>([])
  const [status, setStatus] = React.useState<EventSourceStatus>('closed')

  React.useEffect(() => {
    const es = new EventSource(`${API_ROOT}/v2/opml/status/${jobID}`)
    es.onerror = () => setStatus('error')
    es.onopen = () => setStatus('open')
    es.onmessage = ev => {
      const line = ev.data as string
      if (line === "done") {
        es.close()
        setStatus('closed')
        return
      }
      const parts = line.split('|')
      const jobResultItem:JobResultItem = {
        XMLURL: parts[0],
        error: parts[1] === 'ok' ? undefined : parts[1]
      }
      setItems(data => [jobResultItem].concat(data))
    }

    return () => es.close()
  }, [jobID])

  return (
    <>
      {status === "error" && <Message message="unable to fetch import status" variant="error" />}
      {status === "open" && <Loader />}
      <List dense>
        {items.map((res, idx) => (
          <ListItem key={`import-result-${idx}`}>
            <ListItemIcon>
              {res.error ? <ErrorIcon style={{ color: red[500] }}/> : <SuccessIcon style={{ color: green[500] }}/>}
            </ListItemIcon>
            <ListItemText
              primary={`[${items.length - idx}] ${res.XMLURL}`}
              secondary={res.error}
            />
          </ListItem>
        ))}
      </List>
    </>
  )
}
