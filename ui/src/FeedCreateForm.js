import React, { Component } from 'react'

import {
  Button,
  Form,
  Header,
  Icon,
  Message,
  Modal,
} from 'semantic-ui-react'

import FeedApi from './api/FeedApi'
import FeedTags from './FeedTags'

class FeedCreateForm extends Component {
  constructor(props) {
    super(props)
    this.state = {
      open: false,
      error: null,
      feedXMLURL: '',
      feedTitle: '',
      feedTags: []
    }
  }

  handleOpen = () => this.setState({ open: true, feedXMLURL: '', feedTags: [], error: null })

  handleClose = () => this.setState({ open: false })

  handleChange = (e, { name, value }) => this.setState({ [name]: value })

  handleSubmit = () => {
    const { feedXMLURL, feedTags, feedTitle } = this.state
    FeedApi.add(feedXMLURL, feedTags, feedTitle)
      .then(
        () => {
          this.handleClose()
          this.triggerOnSuccess()
        },
        (error) => this.setState({ error })
      )
  }

  triggerOnSuccess = () => {
    const { onSuccess } = this.props
    onSuccess()
  }

  get errorMessage() {
    const { error } = this.state
    const label = error ? error.message || error.detail : 'unknown'
    return (
      <Message
        error
        header='Unable to add feed'
        content={label}
      />
    )
  }

  render() {
    const { tagSuggestions } = this.props
    const { open, error, feedXMLURL, feedTitle } = this.state
    return (
      <Modal open={open} onClose={this.handleClose} trigger={
        <Button title='Add new feed' floated='right' icon labelPosition='left' primary size='small' onClick={this.handleOpen}>
          <Icon name='plus' /> Add feed
        </Button>
      }>
        <Header icon='rss' content='Add feed' />
        <Modal.Content>
          <Form error={error !== null} onSubmit={this.handleSubmit}>
            <Form.Input
              type='url'
              name='feedXMLURL'
              label='XML URL'
              placeholder='Feed XML URL'
              value={feedXMLURL}
              required
              autoFocus
              onChange={this.handleChange}
            />
            <Form.Input
              type='text'
              name='feedTitle'
              label='Title'
              placeholder='Feed XML title if empty'
              value={feedTitle}
              onChange={this.handleChange}
            />
            <Form.Field>
              <label>Tags</label>
              <FeedTags
                name='feedTags'
                tags={[]}
                suggestions={tagSuggestions}
                onChange={this.handleChange}
              />
            </Form.Field>
            {this.errorMessage}
          </Form>
        </Modal.Content>
        <Modal.Actions>
          <Button onClick={this.handleClose}>
            <Icon name='cancel' /> Cancel
          </Button>
          <Button color='green' onClick={this.handleSubmit}>
            <Icon name='checkmark' /> Add
          </Button>
        </Modal.Actions>
      </Modal>
    )
  }
}

FeedCreateForm.defaultProps = {
  tagSuggestions: [],
}

export default FeedCreateForm
