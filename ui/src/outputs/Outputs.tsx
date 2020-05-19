import React from 'react'

import Loader from '../common/Loader'
import Message from '../common/Message'
import matchResponse from '../helpers/matchResponse'
import { useAPI, usePageTitle } from '../hooks'
import { Output, Entity } from './Types'
import EntityList from './EntityList'

type OutputsResponse = Output[]

const buildEntitiesFromOutputs = (outputs: OutputsResponse) => {
  return outputs.reduce((acc: Entity[], output) => {
    const { filters, ...rest } = output
    const entity: Entity = { ...rest, type: 'output' }
    acc.push(entity)
    if (filters) {
      const sub = filters.map((filter) => {
        const f: Entity = { ...filter, type: 'filter' }
        f.parentId = entity.id
        return f
      })
      acc.push(...sub)
    }
    return acc
  }, [])
}

export default () => {
  usePageTitle('outputs')
  const [loading, outputs, error] = useAPI<OutputsResponse>('/outputs')

  const render = matchResponse<OutputsResponse>({
    Loading: () => <Loader />,
    Data: (data) => <EntityList entities={buildEntitiesFromOutputs(data)} />,
    Error: (err) => <Message message={`Unable to retrieve outputs: ${err.message}`} variant="error" />,
  })

  return <>{render(loading, outputs, error)}</>
}
