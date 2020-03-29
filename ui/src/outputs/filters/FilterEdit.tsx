import React, { useContext, useState } from 'react'
import { RouteComponentProps } from 'react-router'

import { Typography } from '@material-ui/core'

import Loader from '../../common/Loader'
import Message from '../../common/Message'
import { MessageContext } from '../../context/MessageContext'
import fetchAPI from '../../helpers/fetchAPI'
import matchResponse from '../../helpers/matchResponse'
import { useAPI, usePageTitle } from '../../hooks'
import ConfigForm from '../ConfigForm'
import { descFilter, descOutput } from '../helpers'
import { FilterForm, Output } from '../Types'
import { FilterSpecsContext } from './FilterSpecsContext'

type Props = RouteComponentProps<{
  id: string
  filterId: string
}>

export default ({ match, history }: Props) => {
  const { id, filterId } = match.params
  usePageTitle(`edit filter`)
  
  const [error, setError] = useState<Error | null>(null)
  const { showMessage } = useContext(MessageContext)
  const [loading, output, fetchError] = useAPI<Output>(`/outputs/${id}`)
  const { specs } = useContext(FilterSpecsContext)
  
  function handleBack() {
    history.push('/outputs')
  }
  
  async function handleSave(form: FilterForm) {
    try {
      const res = await fetchAPI(`/outputs/${id}/filters/${filterId}`, null, {
        method: 'PUT',
        body: JSON.stringify(form),
      })
      if (!res.ok) {
        const msg = await res.text()
        throw new Error(msg)
      }
      const data = await res.json()
      const filterDesc = descFilter(data)
      showMessage(<Message variant="success"  message={`${filterDesc} configured`} />)
      history.push('/outputs')
    } catch (err) {
      setError(err)
    }
  }

  const render = matchResponse<Output>({
    Loading: () => <Loader />,
    Data: data => {
      const outputDesc = descOutput(data)
      if (!data.filters || data.filters.length === 0) {
        return <Message message={`No filter found for ${outputDesc}`} variant="error" />
      }
      const filter = data.filters.find(f => f.id === filterId)
      if (!filter) {
        return <Message message={`Filter not found in ${outputDesc}`} variant="error" />
      }
      const spec = specs.find(f => f.name === filter.name)
      if (!spec) {
        return <Message message={`Unable to retrieve filter specifications: ${filter.name}`} variant="error" />
      }
      return (
        <>
          <Typography variant="h5" gutterBottom>Configure filter</Typography>
          { !!error && <Message message={error.message} variant="error" />}
          <ConfigForm onSave={handleSave} onCancel={handleBack} spec={spec} source={filter} />
        </>
      )
    },
    Error: err => <Message message={`Unable to fetch filter: ${err.message}`} variant="error" />
  })

  return (<>{render(loading, output, fetchError)}</>)
}
