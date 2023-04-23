import { PropsWithChildren, createElement, forwardRef } from "react"
import classNames from "classnames"

import s from "./IconButton.module.css"

type ButtonProps = { renderAs?: "button" | null } & React.DetailedHTMLProps<React.ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement>;
type AnchorProps = { renderAs: "a" } & React.DetailedHTMLProps<React.AnchorHTMLAttributes<HTMLAnchorElement>, HTMLAnchorElement>;
type Props = ButtonProps | AnchorProps

export const IconButton = forwardRef(function IconButton(props: PropsWithChildren<Props>, ref) {
  const { renderAs, className, children, ...rest } = props

  return createElement(
    renderAs ?? "button",
    {
      className: classNames(s.button,  className),
      ref,
      ...rest,
    },
    children,
  )
})
