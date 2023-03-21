import classNames from "classnames"
import { createElement, PropsWithChildren } from "react"
import s from "./Button.module.css"

type Props = {
    renderAs?: "button" | "a" 
    className?: string
    onClick?: React.MouseEventHandler<HTMLAnchorElement>,
    disabled?: boolean
}

export function Button({ children, renderAs, className, onClick, disabled }: PropsWithChildren<Props>) {
  return createElement(
    renderAs ?? "button",
    {
      className: classNames(s.Button, className),
      onClick,
      disabled,
    },
    children,
  )
}