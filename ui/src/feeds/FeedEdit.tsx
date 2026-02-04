import { useContext, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'

import { Typography } from '@mui/material'

import Loader from '../common/Loader'
import Message from '../common/Message'
import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'
import matchResponse from '../helpers/matchResponse'
import { useAPI, usePageTitle } from '../hooks'
import FeedConfig from './FeedConfig'
import { Feed, FeedForm } from './Types'
import { DefaultHeaders as headers } from '../common/constants'

export default () => {
  const navigate = useNavigate()
  const params = useParams<{ id: string }>()
  const { id } = params
  usePageTitle(`edit feed #${id}`)

  const [error, setError] = useState<Error | null>(null)
  const { showMessage } = useContext(MessageContext)
  const [loading, feeds, fetchError] = useAPI<Feed>(`/feeds/${id}`)

  function handleBack() {
    setError(null)
    navigate('/feeds')
  }

  async function handleSave(form: FeedForm) {
    try {
      const { title, tags } = form
      const body = JSON.stringify({ title, tags })
      const res = await fetchAPI(`/feeds/${id}`, null, { method: 'PUT', headers, body })
      if (!res.ok) {
        const _err = await res.json()
        throw new Error(_err.detail || res.statusText)
      }
      setError(null)
      const data = (await res.json()) as Feed
      showMessage(`${data.title} feed updated`)
      return navigate('/feeds')
    } catch (err: any) {
      setError(err)
    }
  }

  const render = matchResponse<Feed>({
    Loading: () => <Loader />,
    Data: (data) => (
      <>
        <Typography variant="h5" gutterBottom>
          Configure feed
        </Typography>
        {!!error && <Message text={error.message} variant="error" />}
        <FeedConfig onSave={handleSave} onCancel={handleBack} feed={data} />
      </>
    ),
    Error: (err) => <Message text={`Unable to fetch feed: ${err.message}`} variant="error" />,
  })

  return <>{render(loading, feeds, fetchError)}</>
}
