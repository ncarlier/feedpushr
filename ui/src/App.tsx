import { useState } from 'react'
import { HashRouter as Router, Link } from 'react-router-dom'

import { AppBar, Container, CssBaseline, Divider, Drawer, IconButton, Toolbar, Typography } from '@mui/material'
import { blue, pink } from '@mui/material/colors'
import { createTheme, Theme } from '@mui/material/styles'
import { ChevronLeft as ChevronLeftIcon, Info as AboutIcon, Menu as MenuIcon } from '@mui/icons-material'
import { ThemeProvider } from '@mui/material/styles'

import { AuthNProvider } from './context/AuthenticationContext'
import { ConfigProvider } from './context/ConfigContext'
import { MessageProvider } from './context/MessageContext'
import Menu from './Menu'
import Snackbar from './common/Snackbar'
import Routes from './Routes'

const theme = createTheme({
  colorSchemes: {
    dark: false,
  },
})

const drawerWidth = 240

const styles = {
  root: {
    display: 'flex',
  },
  toolbar: {
    paddingRight: '24px', // keep right padding when drawer closed
  },
  toolbarIcon: (theme: Theme) => ({
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'flex-end',
    padding: '0 8px',
    ...theme.mixins.toolbar,
  }),
  appBar: (theme: Theme) => ({
    zIndex: theme.zIndex.drawer + 1,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
  }),
  appBarShift: (theme: Theme) => ({
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth}px)`,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  }),
  menuButton: {
    marginRight: '36px',
  },
  menuButtonHidden: {
    display: 'none',
  },
  title: {
    flexGrow: 1,
  },
  drawerPaper: (theme: Theme) => ({
    position: 'relative' as const,
    whiteSpace: 'nowrap' as const,
    width: drawerWidth,
    transition: theme.transitions.create('width', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  }),
  drawerPaperClose: (theme: Theme) => ({
    overflowX: 'hidden' as const,
    transition: theme.transitions.create('width', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
    width: theme.spacing(7),
    [theme.breakpoints.up('sm')]: {
      width: theme.spacing(6.5),
    },
  }),
  appBarSpacer: (theme: Theme) => theme.mixins.toolbar,
  content: {
    flexGrow: 1,
    height: '100vh',
    overflow: 'auto',
  },
  container: (theme: Theme) => ({
    paddingTop: theme.spacing(4),
    paddingBottom: theme.spacing(4),
  }),
  paper: (theme: Theme) => ({
    padding: theme.spacing(2),
    display: 'flex',
    overflow: 'auto',
    flexDirection: 'column' as const,
  }),
  fixedHeight: {
    height: 240,
  },
}

export default () => {
  const [open, setOpen] = useState(true)
  const handleDrawerOpen = () => setOpen(true)
  const handleDrawerClose = () => setOpen(false)

  return (
    <ConfigProvider>
      <AuthNProvider>
        <Router>
          <ThemeProvider theme={theme}>
            <CssBaseline />
            <AppBar 
              position="absolute" 
              sx={{
                ...styles.appBar(theme),
                ...(open && styles.appBarShift(theme)),
              }}
            >
              <Toolbar sx={styles.toolbar}>
                <IconButton
                  edge="start"
                  color="inherit"
                  aria-label="Open drawer"
                  onClick={handleDrawerOpen}
                  sx={{
                    ...styles.menuButton,
                    ...(open && styles.menuButtonHidden),
                  }}
                >
                  <MenuIcon />
                </IconButton>
                <Typography component="h1" variant="h6" color="inherit" noWrap sx={styles.title}>
                  Feedpushr
                </Typography>
                <IconButton color="inherit" component={Link} to="/about">
                  <AboutIcon />
                </IconButton>
              </Toolbar>
            </AppBar>
            <Drawer
              variant="permanent"
              sx={{
                '& .MuiDrawer-paper': {
                  ...styles.drawerPaper(theme),
                  ...(!open && styles.drawerPaperClose(theme)),
                },
              }}
              open={open}
            >
              <div style={{
                ...styles.toolbarIcon(theme) as React.CSSProperties,
              }}>
                <IconButton onClick={handleDrawerClose}>
                  <ChevronLeftIcon />
                </IconButton>
              </div>
              <Divider />
              <Menu />
            </Drawer>
            <MessageProvider>
              <main style={styles.content}>
                <div style={styles.appBarSpacer(theme) as React.CSSProperties} />
                <Container maxWidth="lg" sx={styles.container(theme)}>
                  <Routes />
                </Container>
              </main>
              <Snackbar />
            </MessageProvider>
          </ThemeProvider>
        </Router>
      </AuthNProvider>
    </ConfigProvider>
  )
}
