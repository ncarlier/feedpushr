import { useContext, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'

import { Typography } from '@mui/material'

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
import { DefaultHeaders as headers } from '../../common/constants'

export default () => {
  const navigate = useNavigate()
  const params = useParams<{ id: string; filterId: string }>()
  const { id, filterId } = params
  usePageTitle(`edit filter`)

  const [error, setError] = useState<Error | null>(null)
  const { showMessage } = useContext(MessageContext)
  const [loading, output, fetchError] = useAPI<Output>(`/outputs/${id}`)
  const { specs } = useContext(FilterSpecsContext)

  function handleBack() {
    navigate('/outputs')
  }

  async function handleSave(form: FilterForm) {
    try {
      const res = await fetchAPI(`/outputs/${id}/filters/${filterId}`, null, {
        method: 'PUT',
        headers,
        body: JSON.stringify(form),
      })
      if (!res.ok) {
        const msg = await res.text()
        throw new Error(msg)
      }
      const data = await res.json()
      const filterDesc = descFilter(data)
      showMessage(`${filterDesc} configured`)
      navigate('/outputs')
    } catch (err: any) {
      setError(err)
    }
  }

  const render = matchResponse<Output>({
    Loading: () => <Loader />,
    Data: (data) => {
      const outputDesc = descOutput(data)
      if (!data.filters || data.filters.length === 0) {
        return <Message text={`No filter found for ${outputDesc}`} variant="error" />
      }
      const filter = data.filters.find((f) => f.id === filterId)
      if (!filter) {
        return <Message text={`Filter not found in ${outputDesc}`} variant="error" />
      }
      const spec = specs.find((f) => f.name === filter.name)
      if (!spec) {
        return <Message text={`Unable to retrieve filter specifications: ${filter.name}`} variant="error" />
      }
      return (
        <>
          <Typography variant="h5" gutterBottom>
            Configure filter
          </Typography>
          {!!error && <Message text={error.message} variant="error" />}
          <ConfigForm onSave={handleSave} onCancel={handleBack} spec={spec} source={filter} />
        </>
      )
    },
    Error: (err) => <Message text={`Unable to fetch filter: ${err.message}`} variant="error" />,
  })

  return <>{render(loading, output, fetchError)}</>
}
