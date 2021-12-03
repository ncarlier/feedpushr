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
  })
)

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
  const classes = useStyles()

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
    <Paper className={classes.root}>
      <form>
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
