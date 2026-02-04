import React, { useCallback } from 'react'

import { Box, Button, Paper, TextField } from '@mui/material'

import { Feed, FeedForm } from './Types'

interface Props {
  feed?: Feed
  onCancel: () => void
  onSave: (feed: FeedForm) => void
}

interface FeedConfigForm {
  title: string
  xmlUrl: string
  tags: string[]
}

export default ({ onSave, onCancel, feed }: Props) => {
  const [values, setValues] = React.useState<FeedConfigForm>({
    title: feed ? feed.title : '',
    xmlUrl: feed ? feed.xmlUrl : '',
    tags: feed && feed.tags ? feed.tags : [],
  })

  const handleChange = (prop: keyof FeedConfigForm) => (event: React.ChangeEvent<HTMLInputElement>) => {
    if (prop === 'tags') {
      setValues({ ...values, [prop]: event.target.value.split(',') })
    } else {
      setValues({ ...values, [prop]: event.target.value })
    }
  }

  const handleSave = useCallback(() => {
    onSave(values)
  }, [onSave, values])

  return (
    <Paper sx={{ padding: 2 }}>
      <Box component="form" sx={{ '& .MuiTextField-root': { m: 1 } }} >
        <TextField id="title" label="Title" value={values.title} onChange={handleChange('title')} fullWidth />
        {feed === undefined && (
          <TextField
          id="xmlurl"
          label="URL"
          type="url"
          helperText="ex: http://rss.cnn.com/rss/edition"
          value={values.xmlUrl}
          onChange={handleChange('xmlUrl')}
          fullWidth
          />
        )}
        <TextField
          id="tags"
          label="Tags"
          helperText="Comma separated list of tags"
          value={values.tags.join(',')}
          onChange={handleChange('tags')}
          fullWidth
        />
      </Box>
      <Button variant="outlined" sx={{ marginRight: 1, marginTop: 2 }} onClick={onCancel}>
        Cancel
      </Button>
      <Button variant="contained" color="primary" sx={{ marginRight: 1, marginTop: 2 }} onClick={handleSave}>
        Save
      </Button>
    </Paper>
  )
}
