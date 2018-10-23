import React, { Component } from 'react'
import timeago from 'timeago.js'

import {
  Button,
  Checkbox,
  Dimmer,
  Icon,
  Label,
  Loader,
  Message,
  Modal,
  Popup,
  Segment,
  Table
} from 'semantic-ui-react'

import FeedApi from './api/FeedApi'
import OPMLApi from './api/OPMLApi'
import UploadButton from './UploadButton'
import FeedTags from './FeedTags'
import FeedCreateForm from './FeedCreateForm'
import FeedUpdateForm from './FeedUpdateForm';

class Feeds extends Component {
  constructor(props) {
    super(props)
    this.state = {
      error: null,
      isLoaded: false,
      items: [],
      column: null,
      direction: null,
    }
  }

  componentDidMount() {
    this.handleRefresh()
  }

  handleToggleAggregation = (id, status) => {
    const action = status ? FeedApi.start(id) : FeedApi.stop(id)
    action.then(
      () => this.handleRefresh(),
      (error) => this.setState({ error })
    )
  }

  handleRemoveFeed = (id) => {
    FeedApi.remove(id)
      .then(
        () => this.handleRefresh(),
        (error) => this.setState({ error })
      )
  }

  handleImportOPML = (file) => {
    OPMLApi.import(file)
      .then(
        (result) => {
          this.setState({ error: null });
          this.handleRefresh()
        },
        (error) => this.setState({ error })
      )
  }

  handleExportOPML = () => {
    OPMLApi.export()
      .then(
        (result) => {
          this.setState({ error: null });
          const element = document.createElement('a')
          element.setAttribute('href', 'data:text/xml;charset=utf-8,' + encodeURIComponent(result))
          element.setAttribute('download', 'export.opml')
          element.style.display = 'none'
          document.body.appendChild(element)
          element.click();
          document.body.removeChild(element);
        },
        (error) => this.setState({ error })
      )
  }

  handleRefresh = () => {
    FeedApi.list()
      .then(
        (result) => {
          let allTags = result.reduce((acc, feed) => {
            if (feed.tags) {
              acc = acc.concat(feed.tags)
            }
            return acc
          }, [])
          allTags = [...(new Set(allTags))].map(tag => ({ id: tag, text: tag }))
          this.setState({
            isLoaded: true,
            error: null,
            items: result,
            allTags: [...(new Set(allTags))]
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

  handleSort = columnToSort => () => {
    const { column, items, direction } = this.state

    if (column !== columnToSort) {
      items.sort(function(a, b) {
        if (columnToSort === 'status') {
          return a.nbProcessedItems - b.nbProcessedItems
        } else {
          return a[columnToSort].localeCompare(b[columnToSort]);
        }
      })
      this.setState({
        column: columnToSort,
        items,
        direction: 'ascending',
      })

      return
    }
    this.setState({
      items: items.reverse(),
      direction: direction === 'ascending' ? 'descending' : 'ascending',
    })
  }

  get tableHeader() {
    const { column, direction } = this.state
    return (
      <Table.Header>
        <Table.Row>
          <Table.HeaderCell />
          <Table.HeaderCell sorted={column === 'title' ? direction : null} onClick={this.handleSort('title')}>
            Feed
          </Table.HeaderCell>
          <Table.HeaderCell sorted={column === 'status' ? direction : null} onClick={this.handleSort('status')}>
            Status
          </Table.HeaderCell>
          <Table.HeaderCell>Last check</Table.HeaderCell>
          <Table.HeaderCell>Next check</Table.HeaderCell>
          <Table.HeaderCell>Created</Table.HeaderCell>
          <Table.HeaderCell>Updated</Table.HeaderCell>
          <Table.HeaderCell />
        </Table.Row>
      </Table.Header>
    )
  }

  get tableBody() {
    const { items, allTags } = this.state;
    return (
      <Table.Body>
        {items.map(item => (
          <Table.Row key={item.id} error={item.errorCount > 0}>
            <Table.Cell collapsing>
              <Checkbox toggle
                defaultChecked={item.status === 'running'}
                title='Start/stop aggregation'
                onChange={(evt, data) => this.handleToggleAggregation(item.id, data.checked)}
              />
            </Table.Cell>
            <Table.Cell>
              <a href={item.xmlUrl} target='_blank'>{item.title}</a>
              <FeedTags tags={item.tags} readonly />
            </Table.Cell>
            <Table.Cell textAlign='center' selectable>{this.renderFeedStatus(item)}</Table.Cell>
            <Table.Cell>{this.renderDate(item.lastCheck)}</Table.Cell>
            <Table.Cell>{this.renderDate(item.nextCheck)}</Table.Cell>
            <Table.Cell>{this.renderDate(item.cdate)}</Table.Cell>
            <Table.Cell>{this.renderDate(item.mdate)}</Table.Cell>
            <Table.Cell textAlign='center'>
              <Modal
                trigger={<Button icon negative size='tiny' title='Remove feed'><Icon name='remove' /></Button>}
                header='Remove feed?'
                content={<Modal.Content>Are you sure to remove feed: <b>{item.title}</b></Modal.Content>}
                actions={['No', { key: `delete-${item.id}`, content: 'Yes', positive: true, onClick: () => this.handleRemoveFeed(item.id) }]}
              />
              <FeedUpdateForm feed={item} tagSuggestions={allTags} onSuccess={this.handleRefresh} />
            </Table.Cell>
          </Table.Row>
        ))}
      </Table.Body>
    )
  }

  get tableFooter() {
    const { allTags } = this.state
    return (
      <Table.Footer fullWidth>
        <Table.Row>
          <Table.HeaderCell />
          <Table.HeaderCell colSpan='7'>
            <FeedCreateForm onSuccess={this.handleRefresh} tagSuggestions={allTags}/>
            <UploadButton size='small' icon onSelectFile={this.handleImportOPML}><Icon name='upload' /> Import</UploadButton>
            <Button size='small' icon onClick={this.handleExportOPML}><Icon name='download' /> Export</Button>
          </Table.HeaderCell>
        </Table.Row>
      </Table.Footer >
    )
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

  get table() {
    const { isLoaded } = this.state;
    return (
      <Segment>
        <Dimmer active={!isLoaded} inverted>
          <Loader inverted>Loading</Loader>
        </Dimmer>
        <Table compact celled definition sortable>
          {this.tableHeader}
          {this.tableBody}
          {this.tableFooter}
        </Table>
      </Segment>
    )
  }

  renderFeedStatus(feed) {
    if (feed.errorCount) {
      return (
        <Popup trigger={<Label circular color='red'>{feed.errorCount}</Label>}>
          {feed.errorMsg}
        </Popup>
      )
    } else if (typeof feed.hubUrl !== 'undefined') {
      const $icon = (<Icon.Group size='large'>
        <Label circular color='green'>{feed.nbProcessedItems}</Label>
        <Icon corner name='cloud' color='green' />
      </Icon.Group>)
      return (
        <Popup trigger={<a href={feed.hubUrl} target='_blank'>{$icon}</a>}>
          PubSubHubbub enabled
        </Popup>
      )
    } else {
      return (
        <Label circular color='green'>{feed.nbProcessedItems}</Label>
      )
    }
  }

  renderDate(date) {
    if (!date) {
      return <span>N/A</span>
    }
    const d = timeago().format(date)
    return (
      <Popup trigger={<span>{d}</span>} content={date} />
    )
  }

  render() {
    return (
      <div>
        {this.errorMessage}
        {this.table}
      </div>
    )
  }
}

export default Feeds
