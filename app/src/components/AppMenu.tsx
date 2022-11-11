import React from "react"
import Grid from "@react-css/grid"
import { NavLink } from "react-router-dom"

const AppMenu = () => {
  return (
    <Grid className="menu">
      <nav>
        <NavLink to="/"><button className="nav-btn">Home</button></NavLink>
        <NavLink to="/movies"><button className="nav-btn">Movies</button></NavLink>
        <NavLink to="/categories"><button className="nav-btn">Categories</button></NavLink>
        <hr />
        <NavLink to="/admin"><button className="nav-btn">Admin</button></NavLink>
      </nav>
    </Grid>
  )
}

export default AppMenu
