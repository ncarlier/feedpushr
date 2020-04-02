import React, { useState, useContext } from 'react'

import {
    Link
} from '@material-ui/core'

import fetchAPI from '../helpers/fetchAPI'
import { usePageTitle } from '../hooks'
import { SearchResult } from './Types'
import MaterialTable, { Query, QueryResult } from 'material-table'
import { MessageContext } from '../context/MessageContext'
import { Feed } from '../feeds/Types'
import Message from '../common/Message'

const headers = {
  "Content-Type": "application/x-www-form-urlencoded",
}

const columns = [
  { 
    title: 'Feed',
    field: 'title',
    render: (feed: SearchResult) => (
      <>
        <Link href={feed.htmlUrl} target="_blank" rel="noreferrer">{feed.title}</Link>
        <p>{feed.desc}</p>
      </>
    )
  },
]

export default () => {
  usePageTitle('Find new RSS feed')
  const [error, setError] = useState<Error | null>(null)
  const [loading, setLoading] = useState<boolean>(false)
  const { showMessage } = useContext(MessageContext)
  
  const save = async (form: SearchResult) => {
    setLoading(true)
    try {
      const { title, xmlUrl: url } = form
      const res = await fetchAPI('/feeds', {title, url}, {method: 'POST', headers})
      if (!res.ok) {
        const _err = await res.json()
        throw new Error(_err.detail || res.statusText)
      }
      setError(null)
      const data = await res.json() as Feed
      showMessage(<Message variant="success"  message={`${data.title} feed created`} />)
    } catch (err) {
      setError(err)
    } finally {
      setLoading(false)
    }
  }
  
  const search = async (query: Query<SearchResult>) => {
    const result: QueryResult<SearchResult> = {
      data: [],
      page: 0,
      totalCount: 0
    }
    const q = query.search.trim()
    if (q.length <= 2) {
      return result
    }
    const res = await fetchAPI('/explore', {q}, {method: 'GET'})
    if (!res.ok) {
      const _err = await res.json()
      throw new Error(_err.detail || res.statusText)
    }
    const data = await res.json() as SearchResult[]
    result.data = data
    result.totalCount = data.length
    return result
  }

  return (
    <>
      { !!error && <Message message={error.message} variant="error" />}
      <MaterialTable
        title="Search"
        columns={ columns }
        data={ search }
        isLoading={ loading }
        options={{
          debounceInterval: 600,
          paging: false,
          pageSize: 10,
          actionsColumnIndex: -1,
        }}
        actions={[
          (rowData: SearchResult) => ({
            icon: 'add_box',
            tooltip: 'Add feed',
            onClick: (event, rowData) => save(rowData as SearchResult),
          })
        ]}
      />
    </>
  )
}
