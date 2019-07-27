import React from 'react'

import {
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
} from '@material-ui/core'

import {
  RssFeed as FeedIcon,
  Transform as FilterIcon,
  Backup as OutputIcon
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
    <ListItem button component={Link} to="/filters">
      <ListItemIcon>
        <FilterIcon />
      </ListItemIcon>
      <ListItemText primary="Filters" />
    </ListItem>
    <ListItem button component={Link} to="/outputs">
      <ListItemIcon>
        <OutputIcon />
      </ListItemIcon>
      <ListItemText primary="Outputs" />
    </ListItem>
  </List>
)
