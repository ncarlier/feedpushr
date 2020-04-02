import React from 'react'

import {
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Divider,
} from '@material-ui/core'

import {
  RssFeed as FeedIcon,
  Backup as OutputIcon,
  Explore as ExploreIcon
} from '@material-ui/icons'
import { Link } from 'react-router-dom';

export default () => (
  <List component="nav" aria-label="Main mailbox folders">
   <ListItem button component={Link} to="/feeds">
      <ListItemIcon>
        <FeedIcon />
      </ListItemIcon>
      <ListItemText primary="Feeds" />
    </ListItem>
    <ListItem button component={Link} to="/outputs">
      <ListItemIcon>
        <OutputIcon />
      </ListItemIcon>
      <ListItemText primary="Outputs" />
    </ListItem>
    <Divider />
    <ListItem button component={Link} to="/explore">
      <ListItemIcon>
        <ExploreIcon />
      </ListItemIcon>
      <ListItemText primary="Explore" />
    </ListItem>
  </List>
)
