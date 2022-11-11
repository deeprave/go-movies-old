import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './components/App'
import { BrowserRouter as Router } from "react-router-dom"

import './index.scss'

const title = "Movie Time!"

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <Router>
      <App title={title} />
    </Router>
  </React.StrictMode>
)
