/*global marked*/
/*eslint no-undef: "error"*/
import React, { useCallback } from 'react'

import { Button, MenuItem, Paper, TextField, Typography } from '@material-ui/core'
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles'

import Doc from '../common/Doc'
import { BaseForm, Filter, Output, Props, Spec } from './Types'

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      padding: theme.spacing(2),
    },
    condition: {
      marginTop: theme.spacing(2),
    },
    button: {
      marginRight: theme.spacing(1),
      marginTop: theme.spacing(2),
    },
  })
)

export type ConfigFormPayload = BaseForm & {
  id?: string
}

interface ConfigFormProps {
  source?: Output | Filter
  spec: Spec
  onCancel: () => void
  onSave: (payload: ConfigFormPayload) => void
}

export default ({ onSave, onCancel, spec, source }: ConfigFormProps) => {
  const classes = useStyles()
  const [alias, setAlias] = React.useState<string>(source ? source.alias : '')
  const [props, setProps] = React.useState<Props>(source ? source.props : {})
  const [condition, setCondition] = React.useState<string>(source ? source.condition : '')

  const handleChangeAlias = useCallback(
    () => (event: React.ChangeEvent<HTMLInputElement>) => {
      setAlias(event.target.value)
    },
    []
  )

  const handleChangeProp = useCallback(
    (name: string) => (event: React.ChangeEvent<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>) => {
      setProps({ ...props, [name]: event.target.value })
    },
    [props]
  )

  const handleChangeCondition = useCallback(
    () => (event: React.ChangeEvent<HTMLInputElement>) => {
      setCondition(event.target.value)
    },
    []
  )

  const handleSave = useCallback(() => {
    onSave({
      alias,
      name: spec.name,
      props,
      condition,
      enabled: source ? source.enabled : false,
    })
  }, [onSave, alias, spec, props, condition, source])

  return (
    <Paper className={classes.root}>
      <Typography variant="h4" gutterBottom>
        {spec.name}
      </Typography>
      <Typography color="textSecondary" dangerouslySetInnerHTML={{ __html: marked(spec.desc) }} />
      <form>
        <Typography variant="h5">Alias</Typography>
        <TextField id="alias" helperText="Alias" value={alias} onChange={handleChangeAlias()} fullWidth />
        {spec.props.length > 0 && <Typography variant="h5">Properties</Typography>}
        {spec.props.map((prop) => (
          <TextField
            id={prop.name}
            key={prop.name}
            label={prop.name}
            helperText={prop.desc}
            value={props[prop.name]}
            type={['select', 'textarea'].includes(prop.type) ? undefined : prop.type}
            multiline={prop.type === 'textarea'}
            select={prop.type === 'select'}
            onChange={handleChangeProp(prop.name)}
            fullWidth
          >
            {prop.options &&
              Object.entries(prop.options).map((option) => (
                <MenuItem key={option[0]} value={option[0]}>
                  {option[1]}
                </MenuItem>
              ))}
          </TextField>
        ))}
        <Typography variant="h5" className={classes.condition}>
          Condition
        </Typography>
        <TextField
          id="condition"
          helperText={
            <>
              Conditional expression (<Doc href="EXPRESSION.md">documentation</Doc>)
            </>
          }
          value={condition}
          onChange={handleChangeCondition()}
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
