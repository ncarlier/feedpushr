
export interface Feed {
  id: string
  title: string
  xmlUrl: string
  hubUrl?: string
  status: string
  tags: string[]
  errorCount?: number
  errorMsg?: string
  nbProcessedItems?: number
  lastCheck?: string
  nextCheck?: string
  cdate: string
  mdate?: string
}

export interface FeedForm {
  title: string
  xmlUrl: string
  tags: string[]
}
