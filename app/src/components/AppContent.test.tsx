import React from "react"
import { expect, describe, it } from 'vitest'
import { render } from "@testing-library/react"
import { BrowserRouter as Router } from "react-router-dom"

import AppContent from "./AppContent"

describe("AppContent component", () => {

  it("renders correctly", () => {
    const appHeader = render(<AppContent />, {wrapper: Router})
    expect(appHeader).to.be.ok
  })

  it("displays a div with class=content", async () => {
    const appHeader = render(<AppContent />, {wrapper: Router})
    const content = appHeader.container as HTMLDivElement
    expect(content.children[0]).toHaveClass("content")
  })

})
