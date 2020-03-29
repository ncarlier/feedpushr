import React, { createContext, ReactNode, useState } from 'react'

import fetchAPI from '../../helpers/fetchAPI'
import { Spec } from '../Types'

interface FilterSpecsContextType {
  specs: Spec[]
}

const FilterSpecsContext = createContext<FilterSpecsContextType>({
  specs: [],
})

interface Props {
  children: ReactNode
}

const FilterSpecsProvider = ({ children }: Props) => {
  const specs = sessionStorage.getItem('filterSpecs') || '[]'
  const [value, setValue] = useState<Spec[]>(JSON.parse(specs))

  const initFilterSpecs = async () => {
    try {
      const res = await fetchAPI('/filters/_specs', null, {method: 'GET'})
      if (!res.ok) {
        throw new Error(res.statusText)
      }
      const data = await res.json()
      sessionStorage.setItem('filterSpecs', JSON.stringify(data))
      setValue(data)
      console.debug('filter specs initialized', data)
    } catch (err) {
      console.error('unable to fetch filters specs', err)
    }
  }

  if (!value.length) {
    initFilterSpecs()
  }

  return (
    <FilterSpecsContext.Provider value={{ specs: value }}>
      {children}
    </FilterSpecsContext.Provider>
  )
}

export { FilterSpecsContext, FilterSpecsProvider }
