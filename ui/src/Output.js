import React, { Component } from 'react'

import {
  Dimmer,
  Item,
  Message,
  Loader,
  Segment,
  Table,
} from 'semantic-ui-react'

import OutputApi from './api/OutputApi'

class Output extends Component {
  constructor(props) {
    super(props)
    this.state = {
      error: null,
      isLoaded: false,
      output: null
    }
  }

  componentDidMount() {
    OutputApi.get()
      .then(
        (result) => {
          this.setState({
            isLoaded: true,
            error: null,
            output: result
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

  get output() {
    const { output } = this.state
    if (output) {
      return (
        <Item.Group>
          <Item>
            <Item.Content>
              <Item.Header>{output.name}</Item.Header>
              <Item.Description>{output.desc}</Item.Description>
              {this.renderOutputProps(output.props)}
            </Item.Content>
          </Item>
        </Item.Group>
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
  }

  renderOutputProps(props) {
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
          {this.output}
        </Segment>
      </div>
    )
  }
}

export default Output