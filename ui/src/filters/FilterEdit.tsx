import React, { useContext, useState } from 'react'
import { RouteComponentProps } from 'react-router'

import { Typography } from '@material-ui/core'

import Loader from '../common/Loader'
import Message from '../common/Message'
import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'
import matchResponse from '../helpers/matchResponse'
import { useAPI, usePageTitle } from '../hooks'
import FilterConfig from './FilterConfig'
import { FilterSpecsContext } from './FilterSpecsContext'
import { Filter, FilterForm } from './Types'

type Props = RouteComponentProps<{id: string}>

export default ({ match, history }: Props) => {
  const { id } = match.params
  usePageTitle(`edit filter #${id}`)
  
  const [error, setError] = useState<Error | null>(null)
  const { showMessage } = useContext(MessageContext)
  const [loading, filters, fetchError] = useAPI<Filter>(`/filters/${id}`)
  const { specs } = useContext(FilterSpecsContext)
  
  function handleBack() {
    history.push('/filters')
  }
  
  async function handleSave(form: FilterForm) {
    try {
      const res = await fetchAPI(`/filters/${id}`, null, {
        method: 'PUT',
        body: JSON.stringify(form),
      })
      if (!res.ok) {
        const msg = await res.text()
        throw new Error(msg)
      }
      const data = await res.json()
      showMessage(<Message variant="success"  message={`Filter ${data.name} (#${data.id}) configured`} />)
      history.push('/filters')
    } catch (err) {
      setError(err)
    }
  }

  const render = matchResponse<Filter>({
    Loading: () => <Loader />,
    Data: data => {
      const spec = specs.find(f => f.name === data.name)
      if (!spec) {
        return <Message message={`Unable to retrieve filter specifications: ${data.name}`} variant="error" />
      }
      return (
        <>
          <Typography variant="h5" gutterBottom>Configure filter</Typography>
          { !!error && <Message message={error.message} variant="error" />}
          <FilterConfig onSave={handleSave} onCancel={handleBack} spec={spec} filter={data} />
        </>
      )
    },
    Error: err => <Message message={`Unable to fetch filter: ${err.message}`} variant="error" />
  })

  return (<>{render(loading, filters, fetchError)}</>)
}
