import React from 'react'

import { Step, StepLabel, Stepper, Typography } from '@material-ui/core'
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles'

import FilterSpecsSelector from './FilterSpecsSelector'

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      width: '90%',
    },
    button: {
      marginRight: theme.spacing(1),
    },
    instructions: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
    },
  }),
)

const steps = ['Select', 'Configure']

export default () => {
  const classes = useStyles()
  const [activeStep, setActiveStep] = React.useState(0)

  function handleNext() {
    setActiveStep(prevActiveStep => prevActiveStep + 1)
  }

  function handleBack() {
    setActiveStep(prevActiveStep => prevActiveStep - 1)
  }

  function handleReset() {
    setActiveStep(0)
  }

  return (
    <div className={classes.root}>
      <Typography variant="h4" gutterBottom>
        Add a filter
      </Typography>
      <Stepper activeStep={activeStep} >
        {steps.map((label, index) => {
          const stepProps: { completed?: boolean } = {}
          const labelProps: { optional?: React.ReactNode } = {}
          return (
            <Step key={label} {...stepProps}>
              <StepLabel {...labelProps}>{label}</StepLabel>
            </Step>
          )
        })}
      </Stepper>
      { activeStep == 0 && <FilterSpecsSelector /> }
      { activeStep == 1 && <p>Configuration...</p> }
    </div>
  )
}
