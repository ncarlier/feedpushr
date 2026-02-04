import { createRoot } from 'react-dom/client'
import App from './App'

// Basic smoke test - renders without crashing
const div = document.createElement('div')
const root = createRoot(div)
root.render(<App />)
root.unmount()
