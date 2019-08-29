import React, { useContext, useState } from 'react'
import { RouteComponentProps, withRouter } from 'react-router'

import { Typography } from '@material-ui/core'

import Message from '../common/Message'
import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'
import { usePageTitle } from '../hooks'
import OutputConfig from './OutputConfig'
import OutputSpecsSelector from './OutputSpecsSelector'
import { OutputForm, OutputSpec } from './Types'

export default withRouter(({ history }: RouteComponentProps) => {
  usePageTitle('add output')
  const [spec, setSpec] = useState<OutputSpec | null>(null)
  const [error, setError] = useState<Error | null>(null)
  const { showMessage } = useContext(MessageContext)

  function handleSelectSpec(spec: OutputSpec) {
    setError(null)
    setSpec(spec)
  }

  function handleBack() {
    setError(null)
    setSpec(null)
  }
  
  async function handleSave(form: OutputForm) {
    try {
      const res = await fetchAPI('/outputs', null, {
        method: 'POST',
        body: JSON.stringify({...form, tags: form.tags.join(',')}),
      })
      if (!res.ok) {
        const msg = await res.text()
        throw new Error(msg)
      }
      const data = await res.json()
      showMessage(<Message variant="success"  message={`Output ${data.name} (#${data.id}) added`} />)
      history.push('/outputs')
    } catch (err) {
      setError(err)
    }
  }

  if (spec === null) {
    return (
      <>
        <Typography variant="h5" gutterBottom>Add output: Choose</Typography>
        <OutputSpecsSelector onSelect={handleSelectSpec} />
      </>
    )
  }

  return (
    <>
      <Typography variant="h5" gutterBottom>Add output: Configure</Typography>
      { !!error && <Message message={error.message} variant="error" />}
      <OutputConfig onSave={handleSave} onCancel={handleBack} spec={spec} />
    </>
  )
})
