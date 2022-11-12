import React from "react"
import { expect, describe, it, vi } from "vitest"
import { render } from "@testing-library/react"

import { Movie } from './Movie'

const mockParam = 1
vi.mock('react-router-dom', () => {
  return {
    useParams: vi.fn(() => ({
      id: mockParam
    }))
  }
})
describe("Movie component", () => {

  it("renders correctly", () => {
    const movie = render(<Movie />)
    expect(movie).to.be.ok
  })

  it("contains a title with the movie id", () => {
    const movie = render(<Movie />)
    const heading = movie.getByRole('heading')
    expect(heading.textContent).toEqual(`Movie:`)
    const table = movie.getByRole('table')
    expect(table).toBeInTheDocument()
  })
})
