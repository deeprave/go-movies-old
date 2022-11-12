import React from "react"
import { expect, describe, it, vi } from "vitest"
import { render } from "@testing-library/react"
import { BrowserRouter } from "react-router-dom"

import { Movies } from "./Movies"

describe("Movies component", () => {

  vi.mock("react-router-dom", async () => ({
    // eslint-disable-next-line
    ...(await vi.importActual<any>("react-router-dom")),
    useLocation() {
      return {
        pathname: '/movies',
        search: '',
        hash: '',
        state: '',
        key: '',
      }
    }
  }))

  it("renders correctly", () => {
    const movies = render(<Movies />, {wrapper: BrowserRouter})
    expect(movies).to.be.ok
  })

  it("component contains a table of movies", () => {
    const movies = render(<Movies />, {wrapper: BrowserRouter})
    const table = movies.queryByRole('table')
    expect(table).to.not.be.null
  })

  it("table contains a thead and tbody", () => {
    const movies = render(<Movies />, {wrapper: BrowserRouter})
    const table = movies.queryByRole('table')
    const thead = table?.querySelector('thead')
    const tbody = table?.querySelector('tbody')
    expect(thead).to.not.be.null
    expect(tbody).to.not.be.null
  })

  it("table contains links to movies", () => {
    const movies = render(<Movies />, {wrapper: BrowserRouter})
    const table = movies.queryByRole('table')
    const links = table?.querySelectorAll('a')
    expect(links).to.not.be.null
    expect(links?.length).toBeGreaterThan(1)
    links?.forEach((link) => {
      expect(link.href).toMatch(/\/movies\/\d+/)
    })
  })

})
