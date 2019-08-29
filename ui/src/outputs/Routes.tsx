import React from 'react'
import { Route, RouteComponentProps, Switch } from 'react-router-dom'

import OutputCreate from './OutputCreate'
import OutputEdit from './OutputEdit'
import Outputs from './Outputs'
import { OutputSpecsProvider } from './OutputSpecsContext'

export default ({ match }: RouteComponentProps) => (
  <OutputSpecsProvider>
    <Switch>
      <Route exact path={match.path + '/'} component={Outputs} />
      <Route exact path={match.path + '/add'} component={OutputCreate} />
      <Route path={match.path + '/:id'} component={OutputEdit} />
    </Switch>
  </OutputSpecsProvider>
)
