import { useEffect } from "react"

import Grid from '@react-css/grid'

import AppHeader from "./AppHeader"
import AppContent from "./AppContent"
import AppFooter from "./AppFooter"
import AppMenu from "./AppMenu"

import './App.scss'

type AppProps = {
  title: string
}

const App = ({title}: AppProps) => {
  useEffect(() => {
    document.title = title
  })

  return (
    <Grid as={'main'} className="App">
      <AppHeader title={title}/>
      <Grid columns="150px auto" gap="1em">
        <AppMenu/>
        <AppContent/>
      </Grid>
      <AppFooter/>
    </Grid>
  )
}

export default App

