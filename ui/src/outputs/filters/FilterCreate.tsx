import { useContext, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'

import { Typography } from '@mui/material'

import Message from '../../common/Message'
import { MessageContext } from '../../context/MessageContext'
import fetchAPI from '../../helpers/fetchAPI'
import { usePageTitle } from '../../hooks'
import ConfigForm from '../ConfigForm'
import SpecSelector from '../SpecSelector'
import { FilterForm, Spec } from '../Types'
import { DefaultHeaders as headers } from '../../common/constants'

export default () => {
  const navigate = useNavigate()
  const params = useParams<{ id: string }>()
  const { id } = params
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
        headers,
        body: JSON.stringify(form),
      })
      if (!res.ok) {
        const msg = await res.text()
        throw new Error(msg)
      }
      const data = await res.json()
      showMessage(`Filter ${data.name} added`)
      navigate('/outputs')
    } catch (err: any) {
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
      {!!error && <Message text={error.message} variant="error" />}
      <ConfigForm onSave={handleSave} onCancel={handleBack} spec={spec} />
    </>
  )
}
