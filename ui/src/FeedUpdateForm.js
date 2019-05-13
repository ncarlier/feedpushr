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

class FeedUpdateForm extends Component {
  constructor(props) {
    super(props)
    const { tags = [] } = props.feed
    this.state = {
      open: false,
      error: null,
      feedTitle: '',
      feedTags: tags
    }
  }

  handleOpen = () => this.setState({ open: true, error: null })

  handleClose = () => this.setState({ open: false })

  handleChange = (e, { name, value }) => this.setState({ [name]: value })

  handleSubmit = () => {
    const { feed } = this.props
    const { feedTags, feedTitle } = this.state
    FeedApi.update(feed, {tags: feedTags, title: feedTitle})
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
    const { feed, tagSuggestions } = this.props
    const { open, error, feedTags, feedTitle } = this.state
    return (
      <Modal open={open} onClose={this.handleClose} trigger={
        <Button title='Update feed' icon size='tiny' onClick={this.handleOpen}>
          <Icon name='edit' />
        </Button>
      }>
        <Header icon='rss' content={`Update feed: ${feed.title}`} />
        <Modal.Content>
          <Form error={error !== null} onSubmit={this.handleSubmit}>
            <Form.Input
              type='text'
              name='feedTitle'
              label='Title'
              placeholder={feed.title}
              value={feedTitle}
              onChange={this.handleChange}
            />
            <Form.Field>
              <label>Tags</label>
              <FeedTags
                name='feedTags'
                tags={ feedTags }
                suggestions={tagSuggestions}
                onChange={this.handleChange}
              />
              <span>Press [enter] or [,] to add a tag</span>
            </Form.Field>
            {this.errorMessage}
          </Form>
        </Modal.Content>
        <Modal.Actions>
          <Button onClick={this.handleClose}>
            <Icon name='cancel' /> Cancel
          </Button>
          <Button color='green' onClick={this.handleSubmit}>
            <Icon name='checkmark' /> Update
          </Button>
        </Modal.Actions>
      </Modal>
    )
  }
}

FeedUpdateForm.defaultProps = {
  tagSuggestions: [],
}

export default FeedUpdateForm
