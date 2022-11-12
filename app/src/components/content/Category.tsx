import React from "react"
import { useParams } from "react-router-dom"

export function Category() {
  const {category} = useParams()
  return <h3>Category: {category}</h3>
}
