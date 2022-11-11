import { expect, describe, it } from "vitest"
import { render } from "@testing-library/react"

import { Home } from './Home'

describe("Home component", () => {

  it("renders correctly", () => {
    const home = render(<Home />)
    expect(home).to.be.ok
  })
})
