import { expect, describe, it } from 'vitest'
import { render } from "@testing-library/react"
import { BrowserRouter as Router } from "react-router-dom"

import AppFooter from "./AppFooter"

describe("AppFooter component", () => {

  it("renders correctly", () => {
    const appFooter = render(<AppFooter />, {wrapper: Router})
    expect(appFooter).to.be.ok
  })

  it("displays the expected footer text", async () => {
    const appFooter = render(<AppFooter />, {wrapper: Router})
    const para = appFooter.getByTestId('footer') as HTMLParagraphElement
    expect(para).toBeInTheDocument()
    expect(para).toHaveTextContent('Copyright')
  })

})
