import React, { ChangeEventHandler, MouseEventHandler, createRef, forwardRef } from 'react'

import { Button, makeStyles } from '@material-ui/core'
import { ButtonProps } from '@material-ui/core/Button'

interface Props {
  onSelectFile: (file: File) => void
}

const useStyles = makeStyles((/*theme: Theme*/) => ({
  hidden: {
    opacity: 0,
    position: 'absolute',
    pointerEvents: 'none',
    width: '1px',
    height: '1px',
  },
}))

export default forwardRef<HTMLButtonElement, Props & ButtonProps>(({ onSelectFile, ...props }, ref) => {
  const classes = useStyles()
  const inputRef = createRef<HTMLInputElement>()

  const handleOnChange: ChangeEventHandler<HTMLInputElement> = (event) => {
    if (event.target.files) {
      const file = event.target.files[0]
      if (file) {
        onSelectFile(file)
      }
    }
  }

  const handleOnClick: MouseEventHandler<HTMLButtonElement> = (/*event*/) => {
    if (inputRef.current) {
      inputRef.current.click()
    }
  }

  return (
    <>
      <input type="file" ref={inputRef} className={classes.hidden} onChange={handleOnChange} />
      <Button {...props} ref={ref} onClick={handleOnClick}>
        {props.children}
      </Button>
    </>
  )
})
