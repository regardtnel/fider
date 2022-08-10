import React from "react"
import { useFider } from "@fider/hooks"
import { classSet } from "@fider/services"

import "./PoweredByFider.scss"

interface PoweredByFiderProps {
  slot: string
  className?: string
}

export const PoweredByFider = (props: PoweredByFiderProps) => {
  // const fider = useFider()

  // const source = encodeURIComponent(fider.session.tenant.subdomain)
  // const medium = "powered-by"
  // const campaign = props.slot

  const className = classSet({
    "c-powered": true,
    [props.className || ""]: props.className,
  })

  return (
    <div className={className}>
      <a rel="noopener" href={`https://costotter.com`} target="_blank">
        Navigate to CostOtter
      </a>
    </div>
  )
}
