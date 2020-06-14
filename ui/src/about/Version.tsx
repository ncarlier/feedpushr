import React from 'react'

import { Typography } from '@material-ui/core'

import Loader from '../common/Loader'
import Message from '../common/Message'
import matchResponse from '../helpers/matchResponse'
import { useAPI } from '../hooks'

interface InfoResponse {
  version: string
}

export default () => {
  const [loading, info, error] = useAPI<InfoResponse>('/')

  const render = matchResponse<InfoResponse>({
    Loading: () => <Loader />,
    Data: (data) => (
      <Typography align="right" color="textSecondary" variant="caption">
        {data.version}
      </Typography>
    ),
    Error: (err) => <Message text={`Unable to retrieve API details: ${err.message}`} variant="error" />,
  })

  return <>{render(loading, info, error)}</>
}
