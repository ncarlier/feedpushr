import MaterialTable, { MTableToolbar } from 'material-table'
import React, { useContext, useState } from 'react'
import { RouteComponentProps, withRouter } from 'react-router'
import { Link } from 'react-router-dom'

import { Link as Href, makeStyles } from '@material-ui/core'
import { Pagination, PaginationItem } from '@material-ui/lab'

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

const useStyles = makeStyles(() => ({
  pagination: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
  },
}))

interface Props {
  page: FeedPage
}

const columns = [
  {
    title: 'Aggregation',
    render: (feed: Feed) => !!feed && <FeedControl feed={feed} />,
    field: 'id',
    sorting: false,
    searchable: false,
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

export default withRouter(({ page, history }: Props & RouteComponentProps) => {
  const classes = useStyles()
  const [error, setError] = useState<Error | null>(null)
  const { showMessage } = useContext(MessageContext)

  const totalPages = Math.ceil(page.total / page.size)

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

  const title = `${page.total} feed${page.total > 1 ? 's' : ''}`

  return (
    <>
      {!!error && <Message message={error.message} variant="error" />}
      <MaterialTable
        title={title}
        columns={columns}
        data={page.data}
        editable={{
          onRowDelete,
        }}
        options={{
          actionsColumnIndex: -1,
          paging: true,
          pageSize: page.data.length,
        }}
        onSearchChange={console.log}
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
          Pagination: () => (
            <td className={classes.pagination}>
              <Pagination
                count={totalPages}
                defaultPage={page.current}
                renderItem={(item: any) => (
                  <PaginationItem
                    component={Link}
                    to={`/feeds${item.page === 1 ? '' : `?page=${item.page}`}`}
                    {...item}
                  />
                )}
              />
            </td>
          ),
        }}
      />
    </>
  )
})
