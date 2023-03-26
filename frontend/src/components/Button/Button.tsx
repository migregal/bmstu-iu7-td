import classNames from "classnames"
import { createElement, PropsWithChildren } from "react"
import s from "./Button.module.css"

type Props = {
    renderAs?: "button" | "a" 
    className?: string
    onClick?: React.MouseEventHandler<HTMLAnchorElement>,
    disabled?: boolean
    outline?: boolean
}

export function Button({ children, renderAs, className, onClick, disabled, outline }: PropsWithChildren<Props>) {
  return createElement(
    renderAs ?? "button",
    {
      className: classNames(s.Button, {[s.Button_outline]: outline },  className),
      onClick,
      disabled,
    },
    children,
  )
}