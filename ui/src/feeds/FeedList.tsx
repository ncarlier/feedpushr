import MaterialTable, { Column, MTableToolbar } from 'material-table'
import React, { useContext, useState } from 'react'

import { Link } from '@material-ui/core'

import Message from '../common/Message'
import Tags from '../common/Tags'
import TimeAgo from '../common/TimeAgo'
import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'
import FeedControl from './FeedControl'
import FeedHub from './FeedHub'
import FeedStatus from './FeedStatus'
import OPMLExportButton from './OPMLExportButton'
import OPMLImportButton from './OPMLImportButton'
import { Feed } from './Types'

const headers = {
  "Content-Type": "application/x-www-form-urlencoded",
}

interface Props {
  feeds: Feed[]
}

const columns: Column[] = [
  { 
    title: 'Aggregation',
    render: (feed: Feed) => ( !!feed && <FeedControl feed={feed} /> ),
    editable: 'never',
    sorting: false,
    searchable: false,
  },
  { 
    title: 'Status',
    field: 'status',
    render: (feed: Feed) => ( !!feed && <FeedStatus feed={feed} /> ),
    editable: 'never',
  },
  { title: 'Title', field: 'title' },
  { 
    title: 'URL',
    field: 'xmlUrl',
    editable: 'onAdd',

    render: (feed: Feed) => (
      <>
        <Link href={feed.xmlUrl} target="_blank">{feed.xmlUrl}</Link>
        &nbsp;
        <FeedHub feed={feed} />
      </>
    )
  },
  { 
    title: 'Tags',
    field: 'tags',
    render: (feed: Feed) => <Tags value={feed.tags} />
  },
  {
    title: 'Next check',
    field: 'nextCheck',
    editable: 'never',
    render: (feed: Feed) => (!!feed && <TimeAgo value={feed.nextCheck} />)
  }
]

export default ({feeds}: Props) => {
  const [data, setData] = useState<Feed[]>(feeds)
  const [error, setError] = useState<Error | null>(null)
  const { showMessage } = useContext(MessageContext)

  const onRowAdd = async (newFeed: Feed) => {
    const { title, xmlUrl: url, tags } = newFeed
    try {
      const res = await fetchAPI('/feeds', {title, url, tags}, {method: 'POST', headers})
      if (res.ok) {
        setError(null)
        const feed = await res.json()
        showMessage(<Message variant="success"  message={`Feed ${feed.title} added`} />)
        return setData([...data, feed])
      }
      const _err = await res.json()
      throw new Error(_err.detail || res.statusText)
    } catch (err) {
      setError(err)
      throw err
    }
  }

  const onRowUpdate = async (newFeed: Feed, oldFeed?: Feed) => {
    const { id, title, tags } = newFeed
    try {
      const res = await fetchAPI(`/feeds/${id}`, {title, tags}, {method: 'PUT', headers})
      if (res.ok) {
        setError(null)
        const feed = await res.json()
        showMessage(<Message variant="success"  message={`Feed ${feed.title} updated`} />)
        return setData(data.map(f => f.id === feed.id ? feed : f))
      }
      const _err = await res.json()
      throw new Error(_err.detail || res.statusText)
    } catch (err) {
      setError(err)
      throw err
    }
  }
  
  const onRowDelete = async (oldFeed: Feed) => {
    const { id, title } = oldFeed
    try {
      const res = await fetchAPI(`/feeds/${id}`, null, {method: 'DELETE'})
      if (res.ok) {
        setError(null)
        showMessage(<Message variant="success"  message={`Feed ${title} removed`} />)
        return setData(data.filter(f => f.id !== id))
      }
      const _err = await res.json()
      throw new Error(_err.detail || res.statusText)
    } catch (err) {
      setError(err)
      throw err
    }
  }

  return <>
    { !!error && <Message message={error.message} variant="error" />}
    <MaterialTable
      title={`${data.length} feeds`}
      columns={ columns }
      data= { data }
      editable = {{
        onRowAdd,
        onRowUpdate,
        onRowDelete
      }}
      options={{
        actionsColumnIndex: -1,
        paging: false
      }}
      components={{
        Toolbar: props => (
          <div>
            <MTableToolbar {...props} />
            <div style={{padding: '5px 10px'}}>
              <OPMLImportButton style={{marginRight: '10px'}}/>
              <OPMLExportButton />
            </div>
          </div>
        ),
      }}
    />
  </>
}
