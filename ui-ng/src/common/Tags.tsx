import React from 'react'

import { Chip } from '@material-ui/core'
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles'

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    tag: {
      marginRight: theme.spacing(0.5),
    },
  }),
)

interface Props {
  value: string[]
}

export default ({value = []}: Props) => (
  <>{ value.map(tag => <Chip key={tag} label={tag} className={useStyles().tag} />) }</>
)
