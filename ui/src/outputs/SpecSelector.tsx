/*global marked*/
/*eslint no-undef: "error"*/

import { useContext } from 'react'

import { Button, Card, CardActions, CardContent, Grid, Typography } from '@mui/material'

import { FilterSpecsContext } from './filters/FilterSpecsContext'
import { OutputSpecsContext } from './OutputSpecsContext'
import { Spec } from './Types'
import { headline } from '../helpers/text'

interface Props {
  type: 'output' | 'filter'
  onSelect: (spec: Spec) => void
}

export default ({ onSelect, type }: Props) => {
  const outputSpecContext = useContext(OutputSpecsContext)
  const filterSpecContext = useContext(FilterSpecsContext)
  const { specs } = type === 'output' ? outputSpecContext : filterSpecContext

  return (
    <Grid spacing={2} container style={{ padding: '0 1em' }}>
      {specs.map((spec) => (
        <Grid item key={spec.name} sm={12} md={4} lg={4}>
          <Card
            sx={{
              heigth: '100%',
              display: 'flex',
              flexDirection: 'column',
            }}
          >
            <CardContent>
              <Typography variant="h5" component="h2">
                {spec.name}
              </Typography>
              <Typography color="textSecondary" dangerouslySetInnerHTML={{ __html: marked(headline(spec.desc)) }} />
            </CardContent>
            <CardActions>
              <Button size="small" onClick={() => onSelect(spec)}>
                Select
              </Button>
            </CardActions>
          </Card>
        </Grid>
      ))}
    </Grid>
  )
}
