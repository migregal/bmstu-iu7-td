import { PropsWithChildren } from "react"

type Props = {
    className?: string
}

export function TableBody(props: PropsWithChildren<Props>) {
  const { children, className } = props
  return <tbody className={className}>{children}</tbody>
}
