import { PropsWithChildren } from "react"
import s from "./Form.module.css"
import classNames from "classnames"

type Props = {
    className?: string
    onSubmit?: (event: React.FormEvent<HTMLFormElement>) => void
    title?: React.ReactNode
    error?: string | null
}

export function Form(props: PropsWithChildren<Props>) {
  const { children, onSubmit, error, className, title } = props

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault()
    event.stopPropagation()
    if (onSubmit) {
      onSubmit(event)
    }
  }

  return <form className={classNames(s.Form, className)} onSubmit={handleSubmit}>
    <h2 className={s.Form__title}>{title}</h2>
    {error ? <p className={classNames(s.Form__field, s.error)}>{error}</p> : null}
    {children}
  </form>
}