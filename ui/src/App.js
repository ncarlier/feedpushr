import React, { Component } from 'react'

import { Route, Redirect } from 'react-router-dom'

import {
  Container,
  Image,
  List,
  Segment,
} from 'semantic-ui-react'

import logo from './logo.svg'

import AppMenu from './AppMenu'
import Feeds from './Feeds'
import Filters from './Filters'
import Output from './Output'

class App extends Component {
  render() {
    return (
      <div>
        <AppMenu />
        <Container style={{ marginTop: '5em' }}>
          <Route exact path="/" render={() => <Redirect to="/feeds"/>} />
          <Route path="/feeds" component={Feeds} />
          <Route path="/filters" component={Filters} />
          <Route path="/output" component={Output} />
        </Container>
        <Segment inverted vertical style={{ margin: '5em 0em 0em', padding: '5em 0em' }}>
          <Container textAlign='center'>
            <Image centered size='mini' src={logo} />
            <List horizontal inverted divided link>
              <List.Item as='a' href='https://github.com/ncarlier/feedpushr'>Sources</List.Item>
              <List.Item as='a' href='https://github.com/ncarlier/feedpushr/issues'>Feature request or bug report</List.Item>
              <List.Item as='a' href='https://www.paypal.me/nunux'>Support project</List.Item>
            </List>
          </Container>
        </Segment>
      </div>
    )
  }
}

export default App
