import React from "react"
import { expect, describe, it } from "vitest"
import { render } from "@testing-library/react"

import { Categories } from "./Categories"

describe("Categories component", () => {

  it("renders correctly", () => {
    const categories = render(<Categories />)
    expect(categories).to.be.ok
  })
})
