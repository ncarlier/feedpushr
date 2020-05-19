import React, { useContext, useState } from 'react'
import { RouteComponentProps } from 'react-router'

import { Typography } from '@material-ui/core'

import Message from '../../common/Message'
import { MessageContext } from '../../context/MessageContext'
import fetchAPI from '../../helpers/fetchAPI'
import { usePageTitle } from '../../hooks'
import ConfigForm from '../ConfigForm'
import SpecSelector from '../SpecSelector'
import { FilterForm, Spec } from '../Types'

type Props = RouteComponentProps<{ id: string }>

export default ({ match, history }: Props) => {
  const { id } = match.params
  usePageTitle('add filter')
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

  async function handleSave(form: FilterForm) {
    try {
      const res = await fetchAPI(`/outputs/${id}/filters`, null, {
        method: 'POST',
        body: JSON.stringify(form),
      })
      if (!res.ok) {
        const msg = await res.text()
        throw new Error(msg)
      }
      const data = await res.json()
      showMessage(<Message variant="success" message={`Filter ${data.name} added`} />)
      history.push('/outputs')
    } catch (err) {
      setError(err)
    }
  }

  if (spec === null) {
    return (
      <>
        <Typography variant="h5" gutterBottom>
          Add filter: Choose
        </Typography>
        <SpecSelector onSelect={handleSelectSpec} type="filter" />
      </>
    )
  }

  return (
    <>
      <Typography variant="h5" gutterBottom>
        Add filter: Configure
      </Typography>
      {!!error && <Message message={error.message} variant="error" />}
      <ConfigForm onSave={handleSave} onCancel={handleBack} spec={spec} />
    </>
  )
}
