import { Chip } from '@mui/material'

import { Feed } from './Types'

export default ({ feed: { tags = [] } }: { feed: Feed }) => {
  return (
    <>
      {tags.map((tag) => (
        <Chip key={tag} label={tag} sx={{ marginRight: 0.5 }} />
      ))}
    </>
  )
}
