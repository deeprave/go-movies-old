import React from "react"
import { Link, useLocation } from "react-router-dom"
import Grid from "@react-css/grid"

export function Categories() {
  const { pathname } = useLocation()

  return (
    <Grid>
      <h3>Categories</h3>
      <table width="100%">
        <thead>
        <tr>
          <th>Category</th>
        </tr>
        </thead>
        <tbody>
        <tr>
          <td>
            <Link to={`${pathname}/comedy`}>Comedy</Link>
          </td>
        </tr>
        <tr>
          <td>
            <Link to={`${pathname}/drama`}>Drama</Link>
          </td>
        </tr>
        </tbody>
      </table>
    </Grid>
  )
}
