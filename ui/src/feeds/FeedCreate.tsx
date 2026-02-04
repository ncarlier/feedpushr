import React, { useContext, useState } from 'react'
import { RouteComponentProps, withRouter } from 'react-router'

import { Typography } from '@material-ui/core'

import Message from '../common/Message'
import { MessageContext } from '../context/MessageContext'
import fetchAPI from '../helpers/fetchAPI'
import { usePageTitle } from '../hooks'
import FeedConfig from './FeedConfig'
import { Feed, FeedForm } from './Types'
import { DefaultHeaders as headers } from '../common/constants'

export default withRouter(({ history }: RouteComponentProps) => {
  usePageTitle('new feed')
  const [error, setError] = useState<Error | null>(null)
  const { showMessage } = useContext(MessageContext)

  function handleBack() {
    setError(null)
    history.push('/feeds')
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
      return history.push('/feeds')
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
})
