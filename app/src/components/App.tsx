import { useEffect } from "react"
import { BrowserRouter as Router } from "react-router-dom"

import AppHeader from "./AppHeader"
import AppContent from "./AppContent"
import AppFooter from "./AppFooter"

import './App.scss'

const App = () => {
  const title = "Movie Time!"
  useEffect(() => {
    document.title = title
  })
  return (
    <div className="App">
      <Router>
        <AppHeader
          title={title}
        />
        <AppContent/>
        <AppFooter/>
      </Router>
    </div>
  )
}

export default App

