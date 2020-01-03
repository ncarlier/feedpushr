import React, { useContext, useState } from 'react'
import { RouteComponentProps, withRouter } from 'react-router'

import { Typography } from '@material-ui/core'

import Message from '../common/Message'
import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'
import { usePageTitle } from '../hooks'
import FilterConfig from './FilterConfig'
import FilterSpecsSelector from './FilterSpecsSelector'
import { FilterForm, FilterSpec } from './Types'

export default withRouter(({ history }: RouteComponentProps) => {
  usePageTitle('add filter')
  const [spec, setSpec] = useState<FilterSpec | null>(null)
  const [error, setError] = useState<Error | null>(null)
  const { showMessage } = useContext(MessageContext)

  function handleSelectSpec(spec: FilterSpec) {
    setError(null)
    setSpec(spec)
  }

  function handleBack() {
    setError(null)
    setSpec(null)
  }
  
  async function handleSave(form: FilterForm) {
    try {
      const res = await fetchAPI('/filters', null, {
        method: 'POST',
        body: JSON.stringify(form),
      })
      if (!res.ok) {
        const msg = await res.text()
        throw new Error(msg)
      }
      const data = await res.json()
      showMessage(<Message variant="success"  message={`Filter ${data.name} (#${data.id}) added`} />)
      history.push('/filters')
    } catch (err) {
      setError(err)
    }
  }

  if (spec === null) {
    return (
      <>
        <Typography variant="h5" gutterBottom>Add filter: Choose</Typography>
        <FilterSpecsSelector onSelect={handleSelectSpec} />
      </>
    )
  }

  return (
    <>
      <Typography variant="h5" gutterBottom>Add filter: Configure</Typography>
      { !!error && <Message message={error.message} variant="error" />}
      <FilterConfig onSave={handleSave} onCancel={handleBack} spec={spec} />
    </>
  )
})
