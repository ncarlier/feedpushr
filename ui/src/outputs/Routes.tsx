import { Route, Routes } from 'react-router-dom'

import OutputCreate from './OutputCreate'
import OutputEdit from './OutputEdit'
import Outputs from './Outputs'
import { OutputSpecsProvider } from './OutputSpecsContext'
import { FilterSpecsProvider } from './filters/FilterSpecsContext'
import FilterCreate from './filters/FilterCreate'
import FilterEdit from './filters/FilterEdit'

export default () => (
  <OutputSpecsProvider>
    <FilterSpecsProvider>
      <Routes>
        <Route path="/" element={<Outputs />} />
        <Route path="/add" element={<OutputCreate />} />
        <Route path="/:id/filters/add" element={<FilterCreate />} />
        <Route path="/:id/filters/:filterId" element={<FilterEdit />} />
        <Route path="/:id" element={<OutputEdit />} />
      </Routes>
    </FilterSpecsProvider>
  </OutputSpecsProvider>
)
