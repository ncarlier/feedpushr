import React, { Component } from 'react'

import {
  Button,
} from 'semantic-ui-react'

class UploadButton extends Component {

  constructor(props) {
    super(props)
    this.styles = {
      opacity: 0,
      position: 'absolute',
      pointerEvents: 'none',
      width: '1px',
      height: '1px'
    }
    this.ref = React.createRef();
  }

  handleChange = (event) => {
    const file = event.target.files[0]
    if (!file) {
      return
    }
    this.props.onSelectFile(file)
  }

  render() {
    const { onSelectFile, ...props } = this.props
    return (
      <div style={{ display: 'inline-block' }}>
        <input type="file" ref={this.ref} style={this.styles} onChange={this.handleChange} />
        <Button {...props} onClick={() => this.ref.current.click()}>{this.props.children}</Button>
      </div>
    )
  }
}

UploadButton.defaultProps = {
  onSelectFile: function () { },
}

export default UploadButton