import { expect, describe, it } from "vitest"
import { render } from "@testing-library/react"
import { BrowserRouter as Router } from "react-router-dom"

import App from "./App"

describe("App top level component", () => {

  it("renders correctly", () => {
    const app = render(<App title=""/>, {wrapper: Router})
    expect(app).to.be.ok
  })

  it("updates document.title", async () => {
    const title = "page title"
    render(<App title={title}/>, {wrapper: Router})
    expect(document.title).toEqual(title)
  })

})
