import React, { Component } from 'react'
import { WithContext as ReactTags } from 'react-tag-input'

import { Label } from 'semantic-ui-react'

const KeyCodes = {
  comma: 188,
  enter: 13,
}

const delimiters = [KeyCodes.comma, KeyCodes.enter]

class FeedTags extends Component {
  constructor (props) {
    super(props)
    this.state = {
      tags: props.tags.map(tag => ({id: tag, text: tag}))
    }
  }

  componentDidUpdate(prevProps) {
    const { tags } = this.props
    if (tags.length !== prevProps.tags.length || !prevProps.tags.every(t => tags.includes(t))) {
      this.setState({
        tags: tags.map(tag => ({ id: tag, text: tag }))
      })
    }
  }

  handleDelete = (i) => {
    const tags = this.state.tags.filter((tag, index) => index !== i)
    this.setState({tags})
    this.triggerOnChange(tags)
  }

  handleAddition = (tag) => {
    const tags = [...this.state.tags, tag]
    this.setState({tags})
    this.triggerOnChange(tags)
  }

  triggerOnChange = (tags) => {
    const { onChange, name } = this.props
    onChange(null, {name, value: tags.map(tag => tag.text)})
  }

  render() {
    const { tags } = this.state
    const { suggestions, readonly = false } = this.props

    if (readonly) {
      return (
        <Label.Group color='blue' size='tiny'>
        { tags.map(tag => (<Label key={`tag-${tag.id}`}>{tag.text}</Label>)) }
        </Label.Group>
      )
    }
    return (
      <ReactTags tags={tags}
        autofocus={false}
        suggestions={suggestions}
        handleDelete={this.handleDelete}
        handleAddition={this.handleAddition}
        delimiters={delimiters}
      />
    )
  }
}

FeedTags.defaultProps = {
  tags: [],
}

export default FeedTags
