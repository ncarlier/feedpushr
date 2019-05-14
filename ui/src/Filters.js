import React, { Component } from 'react'

import {
  Dimmer,
  Item,
  Message,
  Loader,
  Segment,
  Table,
} from 'semantic-ui-react'

import Tags from './Tags'
import FilterApi from './api/FilterApi'

class Filters extends Component {
  constructor(props) {
    super(props)
    this.state = {
      error: null,
      isLoaded: false,
      items: []
    }
  }

  componentDidMount() {
    FilterApi.list()
      .then(
        (result) => {
          this.setState({
            isLoaded: true,
            error: null,
            items: result
          });
        },
        (error) => {
          this.setState({
            isLoaded: true,
            error
          })
        }
      )
  }

  get filters() {
    const { items } = this.state
    if (items.length) {
      return (
        <Item.Group divided>
          {items.map(filter => this.renderFilter(filter))}
        </Item.Group>
      )
    } else {
      return (
        <Message
          warning
          header='No filter loaded!'
          content='You can add filters using the `--filter` daemon parameter.'
        />
      )
    }
  }

  get errorMessage() {
    const { error } = this.state;
    if (error) {
      return (
        <Message negative>
          <Message.Header>An error occured</Message.Header>
          <p>{error.message || error.detail || error.Msg || JSON.stringify(error)}</p>
        </Message>
      )
    }
    return null
  }

  renderFilter(filter) {
    return (
      <Item key={`filter-${filter.name}`}>
        <Item.Content>
          <Item.Header>{filter.name}</Item.Header>
          <Item.Description>
            { filter.tags && <Tags tags={filter.tags} /> }
            <details>
              <summary>Description</summary>
              <pre>{filter.desc}</pre>
            </details>
          </Item.Description>
          {this.renderFilterProps(filter.props)}
        </Item.Content>
      </Item>
    )
  }

  renderFilterProps(props) {
    if (props) {
      return (
        <Item.Extra>
          <Table definition>
            <Table.Body>
              {Object.keys(props).map(prop => (
                <Table.Row key={`prop-${prop}`}>
                  <Table.Cell>{ prop }</Table.Cell>
                  <Table.Cell>{ props[prop] }</Table.Cell>
                </Table.Row>
              ))}
            </Table.Body>
          </Table>
        </Item.Extra>
      )
    }
  }

  render() {
    const { isLoaded } = this.state;
    return (
      <div>
        <Segment>
          <Dimmer active={!isLoaded} inverted>
            <Loader inverted>Loading</Loader>
          </Dimmer>
          {this.errorMessage}
          {this.filters}
        </Segment>
      </div>
    )
  }
}

export default Filters
