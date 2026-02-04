/*eslint no-undef: "error"*/
import React, { useCallback } from 'react'

import { Box, Button, MenuItem, Paper, TextField, Typography } from '@mui/material'

import Doc from '../common/Doc'
import { BaseForm, Filter, Output, Props, Spec } from './Types'
import SpecDesc from './SpecDesc'

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
    <Paper sx={{ padding: 2 }}>
      <Typography variant="h4" gutterBottom>
        {spec.name}
      </Typography>
      <SpecDesc spec={spec} />
      <Box component="form" sx={{ '& .MuiTextField-root': { m: 1 } }} >
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
        <Typography variant="h5" sx={{ marginTop: 2 }}>
          Condition
        </Typography>
        <TextField
          id="condition"
          helperText={
            <>
              Conditional expression (<Doc href="EXPRESSION.md">documentation</Doc>)
            </>
          }
          type="textarea"
          multiline
          value={condition}
          onChange={handleChangeCondition()}
          fullWidth
        />
      </Box>
      <Button variant="outlined" sx={{ marginRight: 1, marginTop: 2 }} onClick={onCancel}>
        Cancel
      </Button>
      <Button variant="contained" color="primary" sx={{ marginRight: 1, marginTop: 2 }} onClick={handleSave}>
        Save
      </Button>
    </Paper>
  )
}
