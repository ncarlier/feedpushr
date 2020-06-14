import MaterialTable from 'material-table'
import React, { useContext, useState } from 'react'
import { RouteComponentProps, withRouter } from 'react-router'

import Ellipsis from '../common/Ellipsis'
import Message from '../common/Message'
import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'
import { Entity } from './Types'
import EntityControl from './EntityControl'
import EntityStatus from './EntityStatus'
import { descEntity } from './helpers'

interface Props {
  entities: Entity[]
}

const columns = [
  {
    title: 'Enabled',
    render: (entity: Entity) => !!entity && <EntityControl entity={entity} />,
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
    render: (entity: Entity) => !!entity && <EntityStatus entity={entity} />,
    searchable: false,
  },
  {
    title: 'Error',
    render: (entity: Entity) => !!entity && <EntityStatus entity={entity} error />,
    searchable: false,
  },
  {
    title: 'Condition',
    field: 'condition',
    render: (entity: Entity) => <Ellipsis value={entity.condition} />,
  },
]

export default withRouter(({ entities, history }: Props & RouteComponentProps) => {
  const [data, setData] = useState<Entity[]>(entities)
  const [error, setError] = useState<Error | null>(null)
  const { showMessage } = useContext(MessageContext)

  const onRowDelete = async (old: Entity) => {
    const { id } = old
    let url = `/outputs/${id}`
    if (old.type === 'filter') {
      url = `/outputs/${old.parentId}/filters/${id}`
    }
    try {
      const res = await fetchAPI(url, null, { method: 'DELETE' })
      if (res.ok) {
        setError(null)
        showMessage(`${descEntity(old)} removed`)
        return setData(data.filter((f) => f.id !== id && f.parentId !== id))
      }
      const _err = await res.json()
      throw new Error(_err.detail || res.statusText)
    } catch (err) {
      setError(err)
      throw err
    }
  }

  return (
    <>
      {!!error && <Message text={error.message} variant="error" />}
      <MaterialTable
        title="Outputs"
        columns={columns}
        data={data}
        parentChildData={(row, rows) => rows.find((a) => a.id === row.parentId)}
        editable={{
          onRowDelete,
        }}
        options={{
          actionsColumnIndex: -1,
          paging: false,
          searchAutoFocus: true,
        }}
        actions={[
          (rowData: Entity) => ({
            icon: 'build',
            tooltip: 'Configure',
            onClick: (event, rowData) => history.push(`/outputs/${(rowData as Entity).id}`),
            hidden: rowData.type !== 'output',
          }),
          (rowData: Entity) => ({
            icon: 'playlist_add',
            tooltip: 'Add filter',
            onClick: (event, rowData) => history.push(`/outputs/${(rowData as Entity).id}/filters/add`),
            hidden: rowData.type !== 'output',
          }),
          (rowData: Entity) => ({
            icon: 'build',
            tooltip: 'Configure',
            onClick: (event, rowData) =>
              history.push(`/outputs/${(rowData as Entity).parentId}/filters/${(rowData as Entity).id}`),
            hidden: rowData.type !== 'filter',
          }),
          {
            icon: 'add_box',
            tooltip: 'Add',
            isFreeAction: true,
            onClick: () => history.push('/outputs/add'),
          },
        ]}
      />
    </>
  )
})
