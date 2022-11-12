import React from "react"
import {BrowserRouter as Router} from "react-router-dom"
import { expect, describe, it } from "vitest"
import { render } from "@testing-library/react"

import { Categories } from "./Categories"

describe("Categories component", () => {

  it("renders correctly", () => {
    const categories = render(<Categories />, {wrapper: Router})
    expect(categories).to.be.ok
  })
})
