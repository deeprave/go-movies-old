import { expect, describe, it } from 'vitest'
import { render } from "@testing-library/react"
import { BrowserRouter as Router } from "react-router-dom"

import AppHeader from "./AppHeader"

describe("AppHeader component", () => {

  it("renders correctly", () => {
    const appHeader = render(<AppHeader title=""/>, {wrapper: Router})
    expect(appHeader).to.be.ok
  })

  it("displays title", async () => {
    const title = "page title"
    const appHeader = render(<AppHeader title={title}/>, {wrapper: Router})
    const heading = appHeader.getByRole('heading')
    expect(heading.textContent).toBe("page title")
  })

})
