import { Tooltip } from '@mui/material'

interface Props {
  value: string
}

const ellips = (value = '', max = 10) => (value.length > max ? value.slice(0, max) + '...' : value)

export default ({ value }: Props) => (
  <Tooltip title={value}>
    <span>{ellips(value)}</span>
  </Tooltip>
)
