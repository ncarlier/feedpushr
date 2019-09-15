import MaterialTable, { Column, MTableToolbar } from 'material-table'
import React, { useContext, useState } from 'react'
import { RouteComponentProps, withRouter } from 'react-router'

import { Link } from '@material-ui/core'

import Message from '../common/Message'
import Tags from '../common/Tags'
import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'
import FeedControl from './FeedControl'
import FeedDates from './FeedDates'
import FeedHub from './FeedHub'
import FeedStatus from './FeedStatus'
import OPMLExportButton from './OPMLExportButton'
import OPMLImportButton from './OPMLImportButton'
import { Feed } from './Types'

interface Props {
  feeds: Feed[]
}

const columns: Column[] = [
  { 
    title: 'Aggregation',
    render: (feed: Feed) => ( !!feed && <FeedControl feed={feed} /> ),
    sorting: false,
    searchable: false,
  },
  { 
    title: 'Status',
    field: 'status',
    render: (feed: Feed) => ( !!feed && <FeedStatus feed={feed} /> ),
  },
  { 
    title: 'Title',
    field: 'title',
    render: (feed: Feed) => (
      <>
        <Link href={feed.xmlUrl} target="_blank">{feed.title}</Link>
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
    render: (feed: Feed) => (!!feed && <FeedDates feed={feed} />)
  }
]

export default withRouter(({feeds, history}: Props & RouteComponentProps) => {
  const [data, setData] = useState<Feed[]>(feeds)
  const [error, setError] = useState<Error | null>(null)
  const { showMessage } = useContext(MessageContext)

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

  let title = data.length > 1 ? `${data.length} feeds` : `${data.length} feed`

  return <>
    { !!error && <Message message={error.message} variant="error" />}
    <MaterialTable
      title={title}
      columns={ columns }
      data= { data }
      editable = {{
        onRowDelete
      }}
      options={{
        actionsColumnIndex: -1,
        paging: false
      }}
      actions={[
        {
          icon: 'edit',
          tooltip: 'Edit',
          onClick: (event, rowData) => history.push(`/feeds/${rowData.id}`)
        },
        {
          icon: 'add_box',
          tooltip: 'Add',
          isFreeAction: true,
          onClick: () => history.push('/feeds/add')
        }
      ]}
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
})
