import React from 'react'

import {
  BottomNavigation,
  BottomNavigationAction,
  createStyles,
  Link,
  makeStyles,
  Paper,
  Theme,
} from '@material-ui/core'
import { Favorite as SupportIcon, NewReleases as FeatureIcon } from '@material-ui/icons'

import logo from './feedpushr.svg'
import SourceIcon from './Github'
import Version from './Version'

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      padding: theme.spacing(3, 2),
      display: 'flex',
      flexDirection: 'column',
    },
    logo: {
      maxWidth: '50vw',
      alignSelf: 'center',
    },
    link: {
      color: 'rgba(0, 0, 0, 0.54)',
    },
  })
)

export default () => {
  const classes = useStyles()

  return (
    <div>
      <Paper className={classes.root}>
        <Version />
        <img src={logo} alt="feedpushr" className={classes.logo} />
        <BottomNavigation showLabels>
          <BottomNavigationAction
            label="Sources"
            icon={<SourceIcon />}
            component={Link}
            href="https://github.com/ncarlier/feedpushr"
            target="_blank"
            rel="noreferrer"
            className={classes.link}
          />
          <BottomNavigationAction
            label="Features &amp; Bugs"
            icon={<FeatureIcon />}
            component={Link}
            href="https://github.com/ncarlier/feedpushr/issues"
            target="_blank"
            rel="noreferrer"
            className={classes.link}
          />
          <BottomNavigationAction
            label="Support this project"
            icon={<SupportIcon />}
            component={Link}
            href="https://www.paypal.me/nunux"
            target="_blank"
            rel="noreferrer"
            className={classes.link}
          />
        </BottomNavigation>
      </Paper>
    </div>
  )
}
