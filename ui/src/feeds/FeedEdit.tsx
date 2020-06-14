import React, { useContext, useState } from 'react'
import { RouteComponentProps } from 'react-router'

import { Typography } from '@material-ui/core'

import Loader from '../common/Loader'
import Message from '../common/Message'
import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'
import matchResponse from '../helpers/matchResponse'
import { useAPI, usePageTitle } from '../hooks'
import FeedConfig from './FeedConfig'
import { Feed, FeedForm } from './Types'

type Props = RouteComponentProps<{ id: string }>

const headers = {
  'Content-Type': 'application/x-www-form-urlencoded',
}

export default ({ match, history }: Props) => {
  const { id } = match.params
  usePageTitle(`edit feed #${id}`)

  const [error, setError] = useState<Error | null>(null)
  const { showMessage } = useContext(MessageContext)
  const [loading, feeds, fetchError] = useAPI<Feed>(`/feeds/${id}`)

  function handleBack() {
    setError(null)
    history.push('/feeds')
  }

  async function handleSave(form: FeedForm) {
    try {
      const { title, tags } = form
      const res = await fetchAPI(`/feeds/${id}`, { title, tags }, { method: 'PUT', headers })
      if (!res.ok) {
        const _err = await res.json()
        throw new Error(_err.detail || res.statusText)
      }
      setError(null)
      const data = (await res.json()) as Feed
      showMessage(`${data.title} feed updated`)
      return history.push('/feeds')
    } catch (err) {
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
