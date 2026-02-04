import {
  BottomNavigation,
  BottomNavigationAction,
  Link,
  Paper,
} from '@mui/material'
import { Favorite as SupportIcon, NewReleases as FeatureIcon } from '@mui/icons-material'

import logo from './feedpushr.svg'
import SourceIcon from './Github'
import Version from './Version'

export default () => {
  return (
    <div>
      <Paper
        sx={{
          padding: 3,
          display: 'flex',
          flexDirection: 'column',
        }}
      >
        <Version />
        <img src={logo} alt="feedpushr" style={{ maxWidth: '50vw', alignSelf: 'center' }} />
        <BottomNavigation showLabels>
          <BottomNavigationAction
            label="Sources"
            icon={<SourceIcon />}
            component={Link}
            href="https://github.com/ncarlier/feedpushr"
            target="_blank"
            rel="noreferrer"
            sx={{ color: 'rgba(0, 0, 0, 0.54)' }}
          />
          <BottomNavigationAction
            label="Features &amp; Bugs"
            icon={<FeatureIcon />}
            component={Link}
            href="https://github.com/ncarlier/feedpushr/issues"
            target="_blank"
            rel="noreferrer"
            sx={{ color: 'rgba(0, 0, 0, 0.54)' }}
          />
          <BottomNavigationAction
            label="Support this project"
            icon={<SupportIcon />}
            component={Link}
            href="https://www.paypal.me/nunux"
            target="_blank"
            rel="noreferrer"
            sx={{ color: 'rgba(0, 0, 0, 0.54)' }}
          />
        </BottomNavigation>
      </Paper>
    </div>
  )
}
