import React from "react"
import { Link, useLocation } from "react-router-dom"
import { IMovie } from "../../common/appTypes"

export function Movies() {
  const [movies, setMovies] = React.useState<IMovie[]>([])
  const {pathname} = useLocation()

  React.useEffect(() => {
    setMovies([
      {id: 1, title: "Movie number one", length: 142},
      {id: 2, title: "Movie number two", length: 137},
      {id: 3, title: "Movie number three", length: 154},
    ])
  }, [])
  return (
    <>
      <h3>Movies</h3>
      <table width="100%">
        <thead>
        <tr>
          <th>Id</th>
          <th>Title</th>
          <th>Length</th>
        </tr>
        </thead>
        <tbody>
        {movies.map((movie) => (
          <tr key={movie.id}>
            <td>{movie.id}</td>
            <td><Link to={`${pathname}/${movie.id}`}>{movie.title}</Link></td>
            <td>{movie.length}</td>
          </tr>
        ))}
        </tbody>
      </table>
    </>
  )
}
