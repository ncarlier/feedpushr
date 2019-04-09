import React from 'react'

import {
  Label,
} from 'semantic-ui-react'

export default ({tags}) => (
  <Label.Group color='blue' size='tiny'>
    { tags.map(tag => (<Label key={`tag-${tag}`}>{tag}</Label>)) }
  </Label.Group>
)
