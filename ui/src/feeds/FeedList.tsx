import MaterialTable, { MTableToolbar, Query, QueryResult } from 'material-table'
import React, { useContext, useState } from 'react'
import { RouteComponentProps, withRouter } from 'react-router'

import { Link as Href } from '@material-ui/core'

import Message from '../common/Message'
import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'
import FeedControl from './FeedControl'
import FeedDates from './FeedDates'
import FeedHtmlLink from './FeedHtmlLink'
import FeedHub from './FeedHub'
import FeedStatus from './FeedStatus'
import FeedTags from './FeedTags'
import OPMLExportButton from './OPMLExportButton'
import OPMLImportButton from './OPMLImportButton'
import { Feed, FeedPage } from './Types'
import { usePageTitle } from '../hooks'

const columns = [
  {
    title: 'Aggregation',
    render: (feed: Feed) => !!feed && <FeedControl feed={feed} />,
    field: 'id',
    width: 120,
  },
  {
    title: 'Status',
    field: 'status',
    render: (feed: Feed) => !!feed && <FeedStatus feed={feed} />,
    width: 100,
  },
  {
    title: 'Title',
    field: 'title',
    render: (feed: Feed) => (
      <>
        <FeedHtmlLink feed={feed} />
        <Href href={feed.xmlUrl} target="_blank">
          {feed.title}
        </Href>
        <FeedHub feed={feed} />
      </>
    ),
  },
  {
    title: 'Tags',
    field: 'tags',
    render: (feed: Feed) => !!feed && <FeedTags feed={feed} />,
  },
  {
    title: 'Next check',
    field: 'nextCheck',
    render: (feed: Feed) => !!feed && <FeedDates feed={feed} />,
  },
]

export default withRouter(({ history }: RouteComponentProps) => {
  usePageTitle('feeds')

  const [error, setError] = useState<Error | null>(null)
  const [title, setTitle] = useState('feeds')
  const { showMessage } = useContext(MessageContext)

  const search = async (query: Query<Feed>): Promise<QueryResult<Feed>> => {
    const req = {
      q: query.search.trim(),
      page: query.page + 1,
      size: query.pageSize,
    }
    const res = await fetchAPI('/feeds', req, { method: 'GET' })
    if (!res.ok) {
      const _err = await res.json()
      throw new Error(_err.detail || res.statusText)
    }
    const page = (await res.json()) as FeedPage
    setTitle(`${page.total} feed${page.total > 1 ? 's' : ''}`)
    return {
      data: page.data,
      page: page.current - 1,
      totalCount: page.total,
    }
  }

  const onRowDelete = async (oldFeed: Feed) => {
    const { id, title } = oldFeed
    try {
      const res = await fetchAPI(`/feeds/${id}`, null, { method: 'DELETE' })
      if (!res.ok) {
        const _err = await res.json()
        throw new Error(_err.detail || res.statusText)
      }
      setError(null)
      showMessage(<Message variant="success" message={`${title} feed removed`} />)
      setTimeout(() => {
        history.push('/')
        history.goBack()
      })
    } catch (err) {
      setError(err)
      throw err
    }
  }

  return (
    <>
      {!!error && <Message message={error.message} variant="error" />}
      <MaterialTable
        title={title}
        columns={columns}
        data={search}
        editable={{
          onRowDelete,
        }}
        options={{
          debounceInterval: 1000,
          actionsColumnIndex: -1,
          paging: true,
          pageSize: 20,
          pageSizeOptions: [10, 20, 50, 100],
          sorting: false,
          searchAutoFocus: true,
        }}
        actions={[
          {
            icon: 'edit',
            tooltip: 'Edit',
            onClick: (event, rowData) => history.push(`/feeds/${(rowData as Feed).id}`),
          },
          {
            icon: 'add_box',
            tooltip: 'Add',
            isFreeAction: true,
            onClick: () => history.push('/feeds/add'),
          },
        ]}
        components={{
          Toolbar: (props) => (
            <div>
              <MTableToolbar {...props} />
              <div style={{ padding: '5px 10px' }}>
                <OPMLImportButton style={{ marginRight: '10px' }} />
                <OPMLExportButton />
              </div>
            </div>
          ),
        }}
      />
    </>
  )
})
