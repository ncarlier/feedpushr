import React, { createContext, ReactNode, useState } from 'react'

import fetchAPI from '../helpers/fetchAPI'
import { Spec } from './Types'

interface OutputSpecsContextType {
  specs: Spec[]
}

const OutputSpecsContext = createContext<OutputSpecsContextType>({
  specs: [],
})

interface Props {
  children: ReactNode
}

const OutputSpecsProvider = ({ children }: Props) => {
  const specs = sessionStorage.getItem('outputSpecs') || '[]'
  const [value, setValue] = useState<Spec[]>(JSON.parse(specs))

  const initOutputSpecs = async () => {
    try {
      const res = await fetchAPI('/outputs/_specs', null, {method: 'GET'})
      if (!res.ok) {
        throw new Error(res.statusText)
      }
      const data = await res.json()
      sessionStorage.setItem('outputSpecs', JSON.stringify(data))
      setValue(data)
      console.debug('output specs initialized', data)
    } catch (err) {
      console.error('unable to fetch outputs specs', err)
    }
  }

  if (!value.length) {
    initOutputSpecs()
  }

  return (
    <OutputSpecsContext.Provider value={{ specs: value }}>
      {children}
    </OutputSpecsContext.Provider>
  )
}

export { OutputSpecsContext, OutputSpecsProvider }
