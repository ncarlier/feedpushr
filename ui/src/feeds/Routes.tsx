import React from 'react'
import { Route, RouteComponentProps, Switch } from 'react-router-dom'

import FeedCreate from './FeedCreate'
import FeedEdit from './FeedEdit'
import FeedList from './FeedList'

export default ({ match }: RouteComponentProps) => (
  <Switch>
    <Route exact path={match.path + '/'} component={FeedList} />
    <Route exact path={match.path + '/add'} component={FeedCreate} />
    <Route path={match.path + '/:id'} component={FeedEdit} />
  </Switch>
)
