import React from "react"
import Grid from "@react-css/grid"
import { Route, Routes } from "react-router-dom"
import { Home } from "./content/Home"
import { Movies } from "./content/Movies"
import { Categories } from "./content/Categories"
import { NoMatch } from "./content/NoMatch"
import { Movie } from "./content/Movie"
import { Category } from "./content/Category"

const AppContent = () => {
  return (
    <Grid className="content">
      <Routes>
        <Route path="/" element={<Home/>} />
        <Route path="/movies/:id" element={<Movie/>} />
        <Route path="/movies" element={<Movies/>} />
        <Route path="/categories/:category" element={<Category/>} />
        <Route path="/categories" element={<Categories/>} />
        <Route path="/*" element={<NoMatch/>} />
      </Routes>
    </Grid>
  )
}

export default AppContent
