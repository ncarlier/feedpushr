import MaterialTable, { Column } from 'material-table'
import React, { useContext, useState } from 'react'
import { RouteComponentProps, withRouter } from 'react-router'

import Message from '../common/Message'
import Tags from '../common/Tags'
import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'
import OutputControl from './OutputControl'
import OutputStatus from './OutputStatus'
import { Output } from './Types'

interface Props {
  outputs: Output[]
}

const columns: Column[] = [
  { 
    title: 'Enabled',
    render: (output: Output) => ( !!output && <OutputControl output={output} /> ),
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
    render: (output: Output) => ( !!output && <OutputStatus output={output} /> ),
    searchable: false,
  },
  { 
    title: 'Error',
    render: (output: Output) => ( !!output && <OutputStatus output={output} error /> ),
    searchable: false,
  },
  { 
    title: 'Tags',
    field: 'tags',
    render: (Output: Output) => <Tags value={Output.tags} />
  }
]

export default withRouter(({outputs, history}: Props & RouteComponentProps) => {
  const [data, setData] = useState<Output[]>(outputs)
  const [error, setError] = useState<Error | null>(null)
  const { showMessage } = useContext(MessageContext)

  const onRowDelete = async (oldOutput: Output) => {
    const { id, name } = oldOutput
    try {
      const res = await fetchAPI(`/outputs/${id}`, null, {method: 'DELETE'})
      if (res.ok) {
        setError(null)
        showMessage(<Message variant="success"  message={`Output ${name} removed`} />)
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
      title="Outputs"
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
          onClick: (event, rowData) => history.push(`/outputs/${rowData.id}`)
        },
        {
          icon: 'add_box',
          tooltip: 'Add',
          isFreeAction: true,
          onClick: () => history.push('/outputs/add')
        }
      ]}
    />
  </>
})
