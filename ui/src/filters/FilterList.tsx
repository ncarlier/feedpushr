import MaterialTable from 'material-table'
import React, { useContext, useState } from 'react'
import { RouteComponentProps, withRouter } from 'react-router'

import Ellipsis from '../common/Ellipsis'
import Message from '../common/Message'
import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'
import FilterControl from './FilterControl'
import FilterStatus from './FilterStatus'
import { Filter } from './Types'

interface Props {
  filters: Filter[]
}

const columns = [
  { 
    title: 'Enabled',
    render: (filter: Filter) => ( !!filter && <FilterControl filter={filter} /> ),
    sorting: false,
    searchable: false,
  },
  { 
    title: 'Alias',
    field: 'alias',
  },
  { 
    title: 'Type',
    field: 'name',
  },
  { 
    title: 'Success',
    render: (filter: Filter) => ( !!filter && <FilterStatus filter={filter} /> ),
    searchable: false,
  },
  { 
    title: 'Error',
    render: (filter: Filter) => ( !!filter && <FilterStatus filter={filter} error /> ),
    searchable: false,
  },
  { 
    title: 'Condition',
    field: 'condition',
    render: (filter: Filter) => <Ellipsis value={filter.condition} />
  }
]

export default withRouter(({filters, history}: Props & RouteComponentProps) => {
  const [data, setData] = useState<Filter[]>(filters)
  const [error, setError] = useState<Error | null>(null)
  const { showMessage } = useContext(MessageContext)

  const onRowDelete = async (oldFilter: Filter) => {
    const { id, name } = oldFilter
    try {
      const res = await fetchAPI(`/filters/${id}`, null, {method: 'DELETE'})
      if (res.ok) {
        setError(null)
        showMessage(<Message variant="success"  message={`Filter ${name} removed`} />)
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
      title="Filters"
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
          icon: 'build',
          tooltip: 'Configure',
          onClick: (event, rowData) => history.push(`/filters/${(rowData as Filter).id}`)
        },
        {
          icon: 'add_box',
          tooltip: 'Add',
          isFreeAction: true,
          onClick: () => history.push('/filters/add')
        }
      ]}
    />
  </>
})
