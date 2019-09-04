/*global marked*/
/*eslint no-undef: "error"*/
import React, { useCallback } from 'react'

import { Button, Paper, TextField, Typography } from '@material-ui/core'
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles'

import { Filter, FilterForm, FilterProps, FilterSpec } from './Types'

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      padding: theme.spacing(2),
    },
    tags: {
      marginTop: theme.spacing(2),
    },
    button: {
      marginRight: theme.spacing(1),
      marginTop: theme.spacing(2),
    },
  }),
)

interface Props {
  filter?: Filter
  spec: FilterSpec
  onCancel: () => void
  onSave: (filter: FilterForm) => void
}

export default ({onSave, onCancel, spec, filter}: Props) => {
  const classes = useStyles()
  const [props, setProps] = React.useState<FilterProps>(filter ? filter.props : {})
  const [tags, setTags] = React.useState<string[]>(filter && filter.tags ? filter.tags : [])

  const handleChangeProp = useCallback((name: string) => (event: React.ChangeEvent<HTMLInputElement>) => {
    setProps({ ...props, [name]: event.target.value })
  }, [props])
  
  const handleChangeTags = useCallback(() => (event: React.ChangeEvent<HTMLInputElement>) => {
    setTags(event.target.value.split(','))
  }, [])

  const handleSave = useCallback(() => {
    onSave({
      name: spec.name,
      props,
      tags,
    })
  }, [onSave, spec, props, tags])

  return (
    <Paper className={classes.root}>
      <Typography variant="h4" gutterBottom>
        {spec.name}
      </Typography>
      <Typography color="textSecondary" dangerouslySetInnerHTML={{__html: marked(spec.desc)}} />
      <form>
        {spec.props.length > 0 && <Typography variant="h5">Properties</Typography>}
        {spec.props.map(prop => (
          <TextField
            id={prop.name}
            key={prop.name}
            label={prop.name}
            helperText={prop.desc}
            type={prop.type}
            defaultValue={props[prop.name]}
            onChange={handleChangeProp(prop.name)}
            fullWidth
          />
        ))}
        <Typography variant="h5" className={classes.tags}>Tags</Typography>
        <TextField
          id="tags"
          helperText="Comma separated list of tags"
          value={tags.join(',')}
          onChange={handleChangeTags()}
          fullWidth
        />
      </form>
      <Button variant="contained" className={classes.button} onClick={onCancel}>
        Cancel
      </Button>
      <Button variant="contained" color="primary" className={classes.button} onClick={handleSave}>
        Save
      </Button>
    </Paper>
  )
}
