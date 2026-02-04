import { Route, Routes } from 'react-router-dom'

import FeedCreate from './FeedCreate'
import FeedEdit from './FeedEdit'
import FeedList from './FeedList'

export default () => (
  <Routes>
    <Route path="/" element={<FeedList />} />
    <Route path="/add" element={<FeedCreate />} />
    <Route path="/:id" element={<FeedEdit />} />
  </Routes>
)
