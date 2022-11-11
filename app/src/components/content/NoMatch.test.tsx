import React from "react"
import { expect, describe, it } from "vitest"
import { render } from "@testing-library/react"
import { NoMatch } from "./NoMatch"

describe("NoMatch component", () => {

  it("renders correctly", () => {
    const nomatch = render(<NoMatch />)
    expect(nomatch).to.be.ok
  })
})
