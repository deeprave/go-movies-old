import React from "react"
import { useParams } from "react-router-dom"

export function Movie() {
  const {id} = useParams()
  return <h3>Movie {id}</h3>
}
