import React, { PropsWithChildren } from "react"
import classNames from "classnames"

import s from "./Placeholder.module.css"

type Props = {
    className?: string
    shouldShow?: boolean
    message?: React.ReactNode
}

export const Placeholder = React.memo(function InnerPlaceholder({ className, message, shouldShow, children }: PropsWithChildren<Props>) {
  return (
    <div className={classNames(s.placeholder__wrapper, className)}>
      <div className={classNames(s.placeholder, shouldShow && s.placeholder_show)}>
        <div className={s.placeholder__text}>{message}</div>
        {children}
      </div>
    </div>
  )
})
