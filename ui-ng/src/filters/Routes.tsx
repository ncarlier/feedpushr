import React from 'react'
import { Route, RouteComponentProps, Switch } from 'react-router-dom'

import FilterCreate from './FilterCreate'
import FilterEdit from './FilterEdit'
import Filters from './Filters'
import { FilterSpecsProvider } from './FilterSpecsContext'

export default ({ match }: RouteComponentProps) => (
  <FilterSpecsProvider>
    <Switch>
      <Route exact path={match.path + '/'} component={Filters} />
      <Route exact path={match.path + '/add'} component={FilterCreate} />
      <Route path={match.path + '/:id'} component={FilterEdit} />
    </Switch>
  </FilterSpecsProvider>
)
