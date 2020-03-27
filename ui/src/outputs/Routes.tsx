import React from 'react'
import { Route, RouteComponentProps, Switch } from 'react-router-dom'

import OutputCreate from './OutputCreate'
import OutputEdit from './OutputEdit'
import Outputs from './Outputs'
import { OutputSpecsProvider } from './OutputSpecsContext'
import { FilterSpecsProvider } from './filters/FilterSpecsContext'
import FilterCreate from './filters/FilterCreate'
import FilterEdit from './filters/FilterEdit'

export default ({ match }: RouteComponentProps) => (
  <OutputSpecsProvider>
    <FilterSpecsProvider>
      <Switch>
        <Route exact path={match.path + '/'} component={Outputs} />
        <Route exact path={match.path + '/add'} component={OutputCreate} />
        <Route exact path={match.path + '/:id/filters/add'} component={FilterCreate} />
        <Route exact path={match.path + '/:id/filters/:filterId'} component={FilterEdit} />
        <Route path={match.path + '/:id'} component={OutputEdit} />
      </Switch>
    </FilterSpecsProvider>
  </OutputSpecsProvider>
)
