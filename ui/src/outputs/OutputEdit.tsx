import { useContext, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'

import { Typography } from '@mui/material'

import Loader from '../common/Loader'
import Message from '../common/Message'
import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'
import matchResponse from '../helpers/matchResponse'
import { useAPI, usePageTitle } from '../hooks'
import ConfigForm from './ConfigForm'
import { OutputSpecsContext } from './OutputSpecsContext'
import { Output, OutputForm } from './Types'
import { DefaultHeaders as headers } from '../common/constants'

export default () => {
  const navigate = useNavigate()
  const params = useParams<{ id: string }>()
  const { id } = params
  usePageTitle('edit output')

  const [error, setError] = useState<Error | null>(null)
  const { showMessage } = useContext(MessageContext)
  const [loading, output, fetchError] = useAPI<Output>(`/outputs/${id}`)
  const { specs } = useContext(OutputSpecsContext)

  function handleBack() {
    navigate('/outputs')
  }

  async function handleSave(form: OutputForm) {
    try {
      const res = await fetchAPI(`/outputs/${id}`, null, {
        method: 'PUT',
        headers,
        body: JSON.stringify(form),
      })
      if (!res.ok) {
        const msg = await res.text()
        throw new Error(msg)
      }
      const data = (await res.json()) as Output
      const desc = data.alias ? data.alias : data.name
      showMessage(`${desc} output configured`)
      navigate('/outputs')
    } catch (err: any) {
      setError(err)
    }
  }

  const render = matchResponse<Output>({
    Loading: () => <Loader />,
    Data: (data) => {
      const spec = specs.find((f) => f.name === data.name)
      if (!spec) {
        return <Message text={`Unable to retrieve output specifications: ${data.name}`} variant="error" />
      }
      return (
        <>
          <Typography variant="h5" gutterBottom>
            Configure output
          </Typography>
          {!!error && <Message text={error.message} variant="error" />}
          <ConfigForm onSave={handleSave} onCancel={handleBack} spec={spec} source={data} />
        </>
      )
    },
    Error: (err) => <Message text={`Unable to fetch output: ${err.message}`} variant="error" />,
  })

  return <>{render(loading, output, fetchError)}</>
}
