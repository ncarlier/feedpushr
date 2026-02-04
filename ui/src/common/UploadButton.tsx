import { ChangeEventHandler, MouseEventHandler, createRef, forwardRef } from 'react'

import { Button, Box } from '@mui/material'
import { ButtonProps } from '@mui/material/Button'

interface Props {
  onSelectFile: (file: File) => void
}

export default forwardRef<HTMLButtonElement, Props & ButtonProps>(({ onSelectFile, ...props }, ref) => {
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
      <Box
        component="input"
        type="file"
        ref={inputRef}
        sx={{
          opacity: 0,
          position: 'absolute',
          pointerEvents: 'none',
          width: '1px',
          height: '1px',
        }}
        onChange={handleOnChange}
      />
      <Button {...props} ref={ref} onClick={handleOnClick}>
        {props.children}
      </Button>
    </>
  )
})
