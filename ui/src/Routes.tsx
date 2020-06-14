import React from 'react'
import { Redirect, Route, Switch } from 'react-router-dom'

import About from './about/About'
import Explore from './explore/Explore'
import FeedRoutes from './feeds/Routes'
import OutputRoutes from './outputs/Routes'

const Routes = () => (
  <Switch>
    <Redirect exact from="/" to="/feeds" />
    <Route path="/feeds" component={FeedRoutes} />
    <Route path="/outputs" component={OutputRoutes} />
    <Route path="/explore" component={Explore} />
    <Route path="/about" component={About} />
  </Switch>
)

export default Routes
