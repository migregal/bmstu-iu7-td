import { PropsWithChildren, createElement } from "react"
import classNames from "classnames"

import s from "./IconButton.module.css"

type ButtonProps = { renderAs?: "button" | null } & React.DetailedHTMLProps<React.ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement>;
type AnchorProps = { renderAs: "a" } & React.DetailedHTMLProps<React.ButtonHTMLAttributes<HTMLAnchorElement>, HTMLAnchorElement>;
type Props = ButtonProps | AnchorProps

export function IconButton(props: PropsWithChildren<Props>) {
  const { renderAs, className, children, ...rest } = props

  return createElement(
    renderAs ?? "button",
    {
      className: classNames(s.button,  className),
      ...rest,
    },
    children,
  )
}
