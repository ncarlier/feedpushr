import { useContext, useState } from 'react'
import { useNavigate } from 'react-router-dom'

import { Typography } from '@mui/material'

import Message from '../common/Message'
import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'
import { usePageTitle } from '../hooks'
import FeedConfig from './FeedConfig'
import { Feed, FeedForm } from './Types'
import { DefaultHeaders as headers } from '../common/constants'

export default () => {
  const navigate = useNavigate()
  usePageTitle('new feed')
  const [error, setError] = useState<Error | null>(null)
  const { showMessage } = useContext(MessageContext)

  function handleBack() {
    setError(null)
    navigate('/feeds')
  }

  async function handleSave(form: FeedForm) {
    try {
      const { title, xmlUrl: url, tags } = form
      const body = JSON.stringify({ title, url, tags })
      const res = await fetchAPI('/feeds', null, { method: 'POST', headers, body })
      if (!res.ok) {
        const _err = await res.json()
        throw new Error(_err.detail || res.statusText)
      }
      setError(null)
      const data = (await res.json()) as Feed
      showMessage(`${data.title} feed created`)
      return navigate('/feeds')
    } catch (err: any) {
      setError(err)
    }
  }

  return (
    <>
      <Typography variant="h5" gutterBottom>
        New feed
      </Typography>
      {!!error && <Message text={error.message} variant="error" />}
      <FeedConfig onSave={handleSave} onCancel={handleBack} />
    </>
  )
}
