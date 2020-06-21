/*global marked*/
import React from 'react'

import { Typography, ExpansionPanel, ExpansionPanelSummary, ExpansionPanelDetails } from '@material-ui/core'
import ExpandMoreIcon from '@material-ui/icons/ExpandMore'

import { headline, afterHeadline } from '../helpers/text'
import { Spec } from './Types'

interface Props {
  spec: Spec
}

export default ({ spec }: Props) => {
  const help = afterHeadline(spec.desc)

  if (help === '') {
    return <Typography color="textSecondary" dangerouslySetInnerHTML={{ __html: marked(headline(spec.desc)) }} />
  }

  return (
    <ExpansionPanel style={{ boxShadow: 'none' }}>
      <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
        <Typography color="textSecondary" dangerouslySetInnerHTML={{ __html: marked(headline(spec.desc)) }} />
      </ExpansionPanelSummary>
      <ExpansionPanelDetails>
        <Typography color="textSecondary" dangerouslySetInnerHTML={{ __html: marked(help) }} />
      </ExpansionPanelDetails>
    </ExpansionPanel>
  )
}
