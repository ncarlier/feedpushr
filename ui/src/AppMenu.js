import React, { Component } from 'react'
import { Link } from 'react-router-dom'

import {
  Container,
  Menu,
  Image
} from 'semantic-ui-react'

import logo from './logo.svg'

export default class AppMenu extends Component {
  state = {}

  handleItemClick = (e, { name }) => this.setState({ activeItem: name })

  render() {
    const { activeItem } = this.state

    return (
      <Menu fixed='top' >
        <Container>
          <Menu.Item>
            <Image
              src={logo}
              size='mini'
              as={Link}
              to='/'
              alt='Feedpushr'
            />
          </Menu.Item>

          <Menu.Item
            name='feeds'
            active={activeItem === 'feeds'}
            as={Link}
            to='/feeds'
            onClick={this.handleItemClick}
          >
            Feeds
        </Menu.Item>

          <Menu.Item
            name='filters'
            active={activeItem === 'filters'}
            as={Link}
            to='/filters'
            onClick={this.handleItemClick}
          >
            Filters
        </Menu.Item>

          <Menu.Item
            name='output'
            active={activeItem === 'output'}
            as={Link}
            to='/output'
            onClick={this.handleItemClick}
          >
            Output
        </Menu.Item>
        </Container>
      </Menu>
    )
  }
}