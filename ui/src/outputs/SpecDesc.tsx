/*global marked*/

import { Typography, Accordion, AccordionSummary, AccordionDetails } from '@mui/material'
import ExpandMoreIcon from '@mui/icons-material/ExpandMore'

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
    <Accordion style={{ boxShadow: 'none' }}>
      <AccordionSummary expandIcon={<ExpandMoreIcon />}>
        <Typography color="textSecondary" dangerouslySetInnerHTML={{ __html: marked(headline(spec.desc)) }} />
      </AccordionSummary>
      <AccordionDetails>
        <Typography color="textSecondary" dangerouslySetInnerHTML={{ __html: marked(help) }} />
      </AccordionDetails>
    </Accordion>
  )
}
