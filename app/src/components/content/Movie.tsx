import React, { useState } from "react"
import { useParams } from "react-router-dom"
import { IMovie } from "../../common/appTypes"
import Grid from "@react-css/grid"

export function Movie() {
  const [movie, setMovie] = useState<IMovie>()
  const {id} = useParams()

  React.useEffect(() => {
    setMovie({id: 0, title: "Some movie", length: 150})
  }, [])

  return (
    <Grid>
      <h3>Movie:</h3>
      <table width="100%">
        <tbody>
          <tr>
            <th>Id</th>
            <td>{id}</td>
          </tr>
          <tr>
            <th>Title</th>
            <td>{movie?.title}</td>
          </tr>
          <tr>
            <th>Length</th>
            <td>{movie?.length}</td>
          </tr>
        </tbody>
      </table>
    </Grid>
  )
}
