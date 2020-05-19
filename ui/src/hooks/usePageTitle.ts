import { useEffect } from 'react'

export default (subtitle?: string, title = 'Feedpushr') => {
  useEffect(() => {
    document.title = subtitle ? `${title} - ${subtitle}` : title
  }, [title, subtitle])
}
