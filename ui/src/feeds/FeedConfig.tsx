import React, { useCallback } from 'react'

import { Button, Paper, TextField } from '@material-ui/core'
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles'

import { Feed, FeedForm } from './Types'

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      padding: theme.spacing(2),
    },
    tags: {
      marginTop: theme.spacing(2),
    },
    button: {
      marginRight: theme.spacing(1),
      marginTop: theme.spacing(2),
    },
  }),
)

interface Props {
  feed?: Feed
  onCancel: () => void
  onSave: (feed: FeedForm) => void
}

export default ({onSave, onCancel, feed}: Props) => {
  const classes = useStyles()
  const [title, setTitle] = React.useState<string>(feed ? feed.title : "")
  const [xmlUrl, setXmlUrl] = React.useState<string>(feed ? feed.xmlUrl : "")
  const [tags, setTags] = React.useState<string[]>(feed && feed.tags ? feed.tags : [])

  const handleChangeTitle = useCallback(() => (event: React.ChangeEvent<HTMLInputElement>) => {
    setTitle(event.target.value)
  }, [])
  const handleChangeXmlUrl = useCallback(() => (event: React.ChangeEvent<HTMLInputElement>) => {
    setXmlUrl(event.target.value)
  }, [])
  const handleChangeTags = useCallback(() => (event: React.ChangeEvent<HTMLInputElement>) => {
    setTags(event.target.value.split(','))
  }, [])

  const handleSave = useCallback(() => {
    onSave({
      title,
      xmlUrl,
      tags,
    })
  }, [onSave, title, xmlUrl, tags])

  return (
    <Paper className={classes.root}>
      <form>
        <TextField
          id="title"
          label="Title"
          value={title}
          onChange={handleChangeTitle()}
          fullWidth
        />
        { !!!feed && <TextField
          id="xmlurl"
          label="URL"
          type="url"
          helperText="ex: http://rss.cnn.com/rss/edition.rss"
          value={xmlUrl}
          onChange={handleChangeXmlUrl()}
          fullWidth
        />}
        <TextField
          id="tags"
          label="Tags"
          helperText="Comma separated list of tags"
          value={tags.join(',')}
          onChange={handleChangeTags()}
          fullWidth
        />
      </form>
      <Button variant="contained" className={classes.button} onClick={onCancel}>
        Cancel
      </Button>
      <Button variant="contained" color="primary" className={classes.button} onClick={handleSave}>
        Save
      </Button>
    </Paper>
  )
}
