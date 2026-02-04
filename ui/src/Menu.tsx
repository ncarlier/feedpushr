import { ListItemIcon, ListItemText, Divider, MenuList, MenuItem } from '@mui/material'

import { RssFeed as FeedIcon, Backup as OutputIcon, Explore as ExploreIcon } from '@mui/icons-material'
import { Link } from 'react-router-dom'

export default () => (
  <MenuList component="nav" aria-label="Main mailbox folders">
    <MenuItem component={Link} to="/feeds">
      <ListItemIcon>
        <FeedIcon />
      </ListItemIcon>
      <ListItemText primary="Feeds" />
    </MenuItem>
    <MenuItem component={Link} to="/outputs">
      <ListItemIcon>
        <OutputIcon />
      </ListItemIcon>
      <ListItemText primary="Outputs" />
    </MenuItem>
    <Divider />
    <MenuItem component={Link} to="/explore">
      <ListItemIcon>
        <ExploreIcon />
      </ListItemIcon>
      <ListItemText primary="Explore" />
    </MenuItem>
  </MenuList>
)
