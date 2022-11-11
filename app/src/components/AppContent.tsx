import Grid from "@react-css/grid"
import { Route, Routes } from "react-router-dom"
import { Home } from "./content/Home"
import { Movies } from "./content/Movies"
import { Categories } from "./content/Categories"
import { NoMatch } from "./content/NoMatch"
import { Movie } from "./content/Movie"

type AppContentProps = {}

const AppContent = ({}: AppContentProps) => {
  return (
    <Grid className="content">
      <Routes>
        <Route path="/" element={<Home/>} />
        <Route path="/movies" element={<Movies/>} />
        <Route path="/movie/:id" element={<Movie/>} />
        <Route path="/categories" element={<Categories/>} />
        <Route path="/*" element={<NoMatch/>} />
      </Routes>
    </Grid>
  )
}

export default AppContent
