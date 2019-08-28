import React from 'react'
import { BrowserRouter as Router, Redirect, Route } from 'react-router-dom'

import {
    AppBar, Badge, Container, CssBaseline, Divider, Drawer, IconButton, Toolbar, Typography
} from '@material-ui/core'
import { makeStyles, Theme } from '@material-ui/core/styles'
import {
    ChevronLeft as ChevronLeftIcon, Menu as MenuIcon, Notifications as NotificationsIcon
} from '@material-ui/icons'

import { MessageProvider } from './context/MessageContext'
import Feeds from './feeds/Feeds'
import FilterRoutes from './filters/Routes'
import classNames from './helpers/classNames'
import Menu from './Menu'
import OutputRoutes from './outputs/Routes'

export default () => {
  const classes = useStyles()

  const [open, setOpen] = React.useState(true)
  const handleDrawerOpen = () => setOpen(true)
  const handleDrawerClose = () => setOpen(false)

  return (
    <Router>
      <CssBaseline />
      <AppBar position="absolute" className={classNames(classes.appBar, open ? classes.appBarShift : null)}>
        <Toolbar className={classes.toolbar}>
          <IconButton
            edge="start"
            color="inherit"
            aria-label="Open drawer"
            onClick={handleDrawerOpen}
            className={classNames(classes.menuButton, open ? classes.menuButtonHidden : null)}
          >
            <MenuIcon />
          </IconButton>
          <Typography component="h1" variant="h6" color="inherit" noWrap className={classes.title}>
            Feedpushr UI
          </Typography>
          <IconButton color="inherit">
            <Badge badgeContent={4} color="secondary">
              <NotificationsIcon />
            </Badge>
          </IconButton>
        </Toolbar>
      </AppBar>
      <Drawer
        variant="permanent"
        classes={{
          paper: classNames(classes.drawerPaper, !open ? classes.drawerPaperClose : null),
        }}
        open={open}
      >
        <div className={classes.toolbarIcon}>
          <IconButton onClick={handleDrawerClose}>
            <ChevronLeftIcon />
          </IconButton>
        </div>
        <Divider />
        <Menu />
      </Drawer>
      <MessageProvider>
        <main className={classes.content}>
          <div className={classes.appBarSpacer} />
          <Container maxWidth="lg" className={classes.container}>
            <Redirect exact from="/" to="/feeds" />
            <Route path="/feeds" component={Feeds} />
            <Route path="/filters" component={FilterRoutes} />
            <Route path="/outputs" component={OutputRoutes} />
          </Container>
        </main>
      </MessageProvider>
    </Router>
  );
}

const drawerWidth = 240

const useStyles = makeStyles<Theme, any>(theme => ({
  root: {
    display: 'flex',
  },
  toolbar: {
    paddingRight: 24, // keep right padding when drawer closed
  },
  toolbarIcon: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'flex-end',
    padding: '0 8px',
    ...theme.mixins.toolbar,
  },
  appBar: {
    zIndex: theme.zIndex.drawer + 1,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
  },
  appBarShift: {
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth}px)`,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  menuButton: {
    marginRight: 36,
  },
  menuButtonHidden: {
    display: 'none',
  },
  title: {
    flexGrow: 1,
  },
  drawerPaper: {
    position: 'relative',
    whiteSpace: 'nowrap',
    width: drawerWidth,
    transition: theme.transitions.create('width', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  drawerPaperClose: {
    overflowX: 'hidden',
    transition: theme.transitions.create('width', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
    width: theme.spacing(7),
    [theme.breakpoints.up('sm')]: {
      width: theme.spacing(9),
    },
  },
  appBarSpacer: theme.mixins.toolbar,
  content: {
    flexGrow: 1,
    height: '100vh',
    overflow: 'auto',
  },
  container: {
    paddingTop: theme.spacing(4),
    paddingBottom: theme.spacing(4),
  },
  paper: {
    padding: theme.spacing(2),
    display: 'flex',
    overflow: 'auto',
    flexDirection: 'column',
  },
  fixedHeight: {
    height: 240,
  },
}))
