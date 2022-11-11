import React from "react"
import Grid from "@react-css/grid"

export type AppHeaderProps = {
  title: string,
}

const AppHeader = ({title}: AppHeaderProps) => {
  return (
    <Grid className="header">
      <h1>{title}</h1>
    </Grid>
  )
}

export default AppHeader
