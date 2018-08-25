import React, { Component } from 'react'
import timeago from 'timeago.js'

import {
    Button,
    Checkbox,
    Dimmer,
    Form,
    Header,
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
import OPMLApi from './api/OPMLApi';
import UploadButton from './UploadButton';

class Feeds extends Component {
    constructor(props) {
        super(props)
        this.state = {
            error: null,
            addError: null,
            isLoaded: false,
            items: [],
            addModalOpen: false,
            feedXMLURL: ''
        }
    }

    componentDidMount() {
        this.handleRefresh()
    }

    handleOpenAddModal = () => this.setState({ addModalOpen: true, feedXMLURL: '', addError: null })

    handleCloseAddModal = () => this.setState({ addModalOpen: false })

    handleChange = (e, { name, value }) => this.setState({ [name]: value })

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

    handleSubmitNewFeed = () => {
        const { feedXMLURL } = this.state
        FeedApi.add(feedXMLURL)
            .then(
                () => {
                    this.handleCloseAddModal()
                    this.handleRefresh()
                },
                (error) => this.setState({ addError: error })
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

    get tableHeader() {
        return (
            <Table.Header>
                <Table.Row>
                    <Table.HeaderCell />
                    <Table.HeaderCell>Title</Table.HeaderCell>
                    <Table.HeaderCell>Status</Table.HeaderCell>
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
        const { items } = this.state;
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
                        <Table.Cell selectable>
                            <a href={item.xmlUrl} target='_blank'>{item.title}</a>
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
                        </Table.Cell>
                    </Table.Row>
                ))}
            </Table.Body>
        )
    }

    get tableFooter() {
        return (
            <Table.Footer fullWidth>
                <Table.Row>
                    <Table.HeaderCell />
                    <Table.HeaderCell colSpan='7'>
                        {this.addFeedForm}
                        <UploadButton size='small' icon onSelectFile={this.handleImportOPML}><Icon name='upload' /> Import</UploadButton>
                        <Button size='small' icon onClick={this.handleExportOPML}><Icon name='download' /> Export</Button>
                    </Table.HeaderCell>
                </Table.Row>
            </Table.Footer>
        )
    }

    get addFeedForm() {
        const { addModalOpen, feedXMLURL, addError } = this.state
        return (
            <Modal open={addModalOpen} onClose={this.handleCloseAddModal} trigger={
                <Button floated='right' icon labelPosition='left' primary size='small' onClick={this.handleOpenAddModal}>
                    <Icon name='plus' /> Add feed
                </Button>
            }>
                <Header icon='rss' content='Add feed' />
                <Modal.Content>
                    <Form error={addError !== null} onSubmit={this.handleSubmitNewFeed}>
                        <Form.Input
                            type='url'
                            name='feedXMLURL'
                            label='XML URL'
                            placeholder='Feed XML URL'
                            value={feedXMLURL}
                            required
                            onChange={this.handleChange}
                        />
                        {this.feedFormError}
                    </Form>
                </Modal.Content>
                <Modal.Actions>
                    <Button onClick={this.handleCloseAddModal}>
                        <Icon name='cancel' /> Cancel
                    </Button>
                    <Button color='green' onClick={this.handleSubmitNewFeed}>
                        <Icon name='checkmark' /> Add
                    </Button>
                </Modal.Actions>
            </Modal>
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
                <Table compact celled definition>
                    {this.tableHeader}
                    {this.tableBody}
                    {this.tableFooter}
                </Table>
            </Segment>
        )
    }

    get feedFormError() {
        const { addError } = this.state
        const label = addError ? addError.message || addError.detail : 'unknown'
        return (
            <Message
                error
                header='Unable to add feed'
                content={label}
            />
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
                <Icon name='checkmark' color='green' />
                <Icon corner name='cloud' color='green' />
            </Icon.Group>)
            return (
                <Popup trigger={<a href={feed.hubUrl} target='_blank'>{$icon}</a>}>
                    PubSubHubbub enabled
                </Popup>
            )
        } else {
            return (<Icon color='green' name='checkmark' size='large' />)
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