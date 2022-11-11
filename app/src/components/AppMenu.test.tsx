import { expect, describe, it } from 'vitest'
import { render } from "@testing-library/react"
import { BrowserRouter as Router } from "react-router-dom"

import AppMenu from "./AppMenu"

describe("AppMenu component", () => {

  it("renders correctly", () => {
    const appMenu = render(<AppMenu />, {wrapper: Router})
    expect(appMenu).to.be.ok
  })

  it("displays div with class=menu", async () => {
    const appMenu = render(<AppMenu />, {wrapper: Router})
    const content = appMenu.container as HTMLDivElement
    expect(content.children[0]).toHaveClass("menu")
  })

  it("displays a menu of links", async () => {
    const appMenu = render(<AppMenu />, {wrapper: Router})
    const links = appMenu.getAllByRole('link')
    expect(links.length).toBeGreaterThan(2)
  })

})
