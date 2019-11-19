/*global marked*/
/*eslint no-undef: "error"*/
import React, { useCallback } from 'react'

import { Button, Paper, TextField, Typography, MenuItem } from '@material-ui/core'
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles'

import { Output, OutputForm, OutputProps, OutputSpec } from './Types'

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
  output?: Output
  spec: OutputSpec
  onCancel: () => void
  onSave: (output: OutputForm) => void
}

export default ({onSave, onCancel, spec, output}: Props) => {
  const classes = useStyles()
  const [alias, setAlias] = React.useState<string>(output ? output.alias : "")
  const [props, setProps] = React.useState<OutputProps>(output ? output.props : {})
  const [tags, setTags] = React.useState<string[]>(output && output.tags ? output.tags : [])

  const handleChangeAlias = useCallback(() => (event: React.ChangeEvent<HTMLInputElement>) => {
    setAlias(event.target.value)
  }, [])

  const handleChangeProp = useCallback((name: string) => (event: React.ChangeEvent<HTMLInputElement|HTMLSelectElement>) => {
    setProps({ ...props, [name]: event.target.value })
  }, [props])
  
  const handleChangeTags = useCallback(() => (event: React.ChangeEvent<HTMLInputElement>) => {
    setTags(event.target.value.split(','))
  }, [])

  const handleSave = useCallback(() => {
    onSave({
      alias,
      name: spec.name,
      props,
      tags,
    })
  }, [onSave, alias, spec, props, tags])

  return (
    <Paper className={classes.root}>
      <Typography variant="h4" gutterBottom>
        {spec.name}
      </Typography>
      <Typography color="textSecondary" dangerouslySetInnerHTML={{__html: marked(spec.desc)}} />
      <form>
        <Typography variant="h5">Alias</Typography>
        <TextField
          id="alias"
          helperText="Alias"
          value={alias}
          onChange={handleChangeAlias()}
          fullWidth
        />
        {spec.props.length > 0 && <Typography variant="h5">Properties</Typography>}
        {spec.props.map(prop => (
          <TextField
            id={prop.name}
            key={prop.name}
            label={prop.name}
            helperText={prop.desc}
            value={props[prop.name]}
            type={['select', 'textarea'].includes(prop.type) ? undefined : prop.type}
            multiline={prop.type === 'textarea'}
            select={prop.type === 'select'}
            defaultValue={props[prop.name]}
            onChange={handleChangeProp(prop.name)}
            fullWidth
          >
            {prop.options && prop.options.map(option => (
              <MenuItem key={option} value={option}>
                {option}
              </MenuItem>
            ))}
          </TextField>
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
