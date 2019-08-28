import React from 'react'

import {
    BottomNavigation, BottomNavigationAction, createStyles, Link, makeStyles, Paper, Theme
} from '@material-ui/core'
import { Favorite as SupportIcon, NewReleases as FeatureIcon } from '@material-ui/icons'

import logo from './feedpushr.svg'
import SourceIcon from './Github'

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      padding: theme.spacing(3, 2),
      textAlign: "center",
    },
    logo: {
      maxWidth: "50vw",
    }
  }),
)

export default () => {
  const classes = useStyles()

  return (
    <div>
      <Paper className={classes.root}>
        <img src={logo} alt="feedpushr" className={classes.logo}/>
        <BottomNavigation showLabels>
          <BottomNavigationAction
            label="Sources"
            icon={<SourceIcon />}
            component={Link}
            href="https://github.com/ncarlier/feedpushr"
            target="_blank"
            rel="noreferrer"
          />
          <BottomNavigationAction
            label="Features &amp; Bugs"
            icon={<FeatureIcon />}
            component={Link}
            href="https://github.com/ncarlier/feedpushr/issues"
            target="_blank"
            rel="noreferrer"
          />
          <BottomNavigationAction
            label="Support this project"
            icon={<SupportIcon />}
            component={Link}
            href="https://www.paypal.me/nunux"
            target="_blank"
            rel="noreferrer"
          />
        </BottomNavigation>
      </Paper>
    </div>
  )
}
