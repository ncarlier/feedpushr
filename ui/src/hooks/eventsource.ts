import { useEffect, useRef, useState } from 'react'

type EventSourceConstructor = {
  new (url: string, eventSourceInitDict?: EventSourceInit): EventSource
}

export type EventSourceStatus = 'init' | 'open' | 'closed' | 'error'

export type EventSourceEvent = Event & { data: string }

export function useEventSource(url: string, withCredentials?: boolean, ESClass: EventSourceConstructor = EventSource) {
  const source = useRef<EventSource | null>(null)
  const [status, setStatus] = useState<EventSourceStatus>('init')
  useEffect(() => {
    if (url) {
      const es = new ESClass(url, { withCredentials })
      source.current = es

      es.addEventListener('open', () => setStatus('open'))
      es.addEventListener('error', () => setStatus('error'))

      return () => {
        source.current = null
        es.close()
      }
    }

    setStatus('closed')

    return undefined
  }, [url, withCredentials, ESClass])

  return [source.current, status] as const
}

export function useEventSourceListener(
  source: EventSource | null,
  types: string[],
  listener: (e: EventSourceEvent) => void,
  dependencies: any[] = []
) {
  useEffect(() => {
    if (source) {
      types.forEach((type) => source.addEventListener(type, listener as any))
      return () => types.forEach((type) => source.removeEventListener(type, listener as any))
    }
    return undefined
  }, [source, listener, types, ...dependencies])
}
