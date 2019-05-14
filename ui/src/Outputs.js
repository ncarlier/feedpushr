import React, { useState, useEffect } from 'react'

import {
  Dimmer,
  Item,
  Message,
  Loader,
  Segment,
  Table,
} from 'semantic-ui-react'

import Tags from './Tags'
import OutputApi from './api/OutputApi'

function Error({err}) {
  if (!err) {
    return null
  }
  return (
    <Message negative>
      <Message.Header>An error occured</Message.Header>
      <p>{err.message || err.detail || err.Msg || JSON.stringify(err)}</p>
    </Message>
  )
}

function OutputProperties({properties}) {
  if (!properties) {
    return null
  }
  return (
    <Item.Extra>
      <Table definition>
        <Table.Body>
          {Object.keys(properties).map(prop => (
            <Table.Row key={`prop-${prop}`}>
              <Table.Cell>{ prop }</Table.Cell>
              <Table.Cell>{ properties[prop] }</Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table>
    </Item.Extra>
  )
}

function OutputItem({data}) {
  if (!data) {
    return null
  }
  
  return (
    <Item.Group>
      <Item>
        <Item.Content>
          <Item.Header>{data.name}</Item.Header>
          <Item.Description>
            { data.tags && <Tags tags={data.tags} /> }
            <details>
              <summary>Description</summary>
              <pre>{data.desc}</pre>
            </details>
          </Item.Description>
          <OutputProperties properties={data.props} />
        </Item.Content>
      </Item>
    </Item.Group>
  )
}

export default () => {
  const [data, setData] = useState(null)
  const [error, setError] = useState(null)
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    setIsLoading(true)
    OutputApi.list()
    .then(
      data => setData(data),
      error => setError(error)
    )
    .then(() => setIsLoading(false))
  }, [])

  return (
    <div>
      <Segment>
        <Dimmer active={isLoading} inverted>
          <Loader inverted>Loading</Loader>
        </Dimmer>
        <Error err={error} />
        { data && data.length &&
          <Item.Group divided>
            {data.map(output => <OutputItem key={`output-${output.name}`} data={output}/>)}
          </Item.Group>
        }
      </Segment>
    </div>
  )
}
