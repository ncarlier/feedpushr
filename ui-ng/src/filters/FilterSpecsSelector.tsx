/*global marked*/
/*eslint no-undef: "error"*/

import React, { useContext } from 'react'

import { Button, Card, CardActions, CardContent, Grid, Typography } from '@material-ui/core'
import { makeStyles } from '@material-ui/core/styles'

import { FilterSpecsContext } from './FilterSpecsContext'

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

export default () => {
  const classes = useStyles()
  const { specs } = useContext(FilterSpecsContext)

  return (
    <Grid spacing={2} container style={{padding: '0 1em'}}>
      { specs.map(spec => (
        <Grid item key={spec.name} sm={12} md={4} lg={4}>
          <Card className={classes.card}>
            <CardContent>
              <Typography variant="h5" component="h2">{spec.name}</Typography>
              <Typography className={classes.pos} color="textSecondary">
                contrib
              </Typography>
              <div dangerouslySetInnerHTML={{__html: marked(spec.desc)}} />
            </CardContent>
            <CardActions>
              <Button size="small">Select</Button>
            </CardActions>
          </Card>
        </Grid>
      ))}
    </Grid>
  )
}
