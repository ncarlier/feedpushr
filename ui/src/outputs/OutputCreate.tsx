import React, { useContext, useState } from 'react'
import { RouteComponentProps, withRouter } from 'react-router'

import { Typography } from '@material-ui/core'

import Message from '../common/Message'
import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'
import { usePageTitle } from '../hooks'
import ConfigForm from './ConfigForm'
import SpecSelector from './SpecSelector'
import { Output, OutputForm, Spec } from './Types'

export default withRouter(({ history }: RouteComponentProps) => {
  usePageTitle('add output')
  const [spec, setSpec] = useState<Spec | null>(null)
  const [error, setError] = useState<Error | null>(null)
  const { showMessage } = useContext(MessageContext)

  function handleSelectSpec(spec: Spec) {
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
        body: JSON.stringify(form),
      })
      if (!res.ok) {
        const msg = await res.text()
        throw new Error(msg)
      }
      const data = (await res.json()) as Output
      const desc = data.alias ? data.alias : data.name
      showMessage(`${desc} output added`)
      history.push('/outputs')
    } catch (err) {
      setError(err)
    }
  }

  if (spec === null) {
    return (
      <>
        <Typography variant="h5" gutterBottom>
          Add output: Choose
        </Typography>
        <SpecSelector onSelect={handleSelectSpec} type="output" />
      </>
    )
  }

  return (
    <>
      <Typography variant="h5" gutterBottom>
        Add output: Configure
      </Typography>
      {!!error && <Message text={error.message} variant="error" />}
      <ConfigForm onSave={handleSave} onCancel={handleBack} spec={spec} />
    </>
  )
})
