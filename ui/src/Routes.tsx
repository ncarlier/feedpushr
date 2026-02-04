import { Navigate, Route, Routes as RouterRoutes } from 'react-router-dom'

import About from './about/About'
import Explore from './explore/Explore'
import FeedRoutes from './feeds/Routes'
import OutputRoutes from './outputs/Routes'

const Routes = () => (
  <RouterRoutes>
    <Route path="/" element={<Navigate to="/feeds" replace />} />
    <Route path="/feeds/*" element={<FeedRoutes />} />
    <Route path="/outputs/*" element={<OutputRoutes />} />
    <Route path="/explore" element={<Explore />} />
    <Route path="/about" element={<About />} />
  </RouterRoutes>
)

export default Routes
