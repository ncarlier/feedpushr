/*global marked*/
/*eslint no-undef: "error"*/

import React, { useContext } from 'react'

import { Button, Card, CardActions, CardContent, Grid, Typography } from '@material-ui/core'
import { makeStyles } from '@material-ui/core/styles'

import excerpt from '../helpers/excerpt'
import { OutputSpecsContext } from './OutputSpecsContext'
import { OutputSpec } from './Types'

const useStyles = makeStyles({
  card: {
    heigth: '100%',
    display: 'flex',
    flexDirection: 'column',
  },
  title: {
    fontSize: 14,
  },
  pos: {
    marginBottom: 12,
  },
})

interface Props {
  onSelect: (spec: OutputSpec) => void
}

export default ({onSelect}: Props) => {
  const classes = useStyles()
  const { specs } = useContext(OutputSpecsContext)

  return (
    <Grid spacing={2} container style={{padding: '0 1em'}}>
      { specs.map(spec => (
        <Grid item key={spec.name} sm={12} md={4} lg={4}>
          <Card className={classes.card}>
            <CardContent>
              <Typography variant="h5" component="h2">{spec.name}</Typography>
              <Typography color="textSecondary" dangerouslySetInnerHTML={{__html: marked(excerpt(spec.desc))}} />
            </CardContent>
            <CardActions>
              <Button size="small" onClick={() => onSelect(spec)}>Select</Button>
            </CardActions>
          </Card>
        </Grid>
      ))}
    </Grid>
  )
}
