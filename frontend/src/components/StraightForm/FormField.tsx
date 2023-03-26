import s from "./Form.module.css"
import classNames from "classnames"

type Props = {
    className?: string
    label?: React.ReactNode,
    type?: React.HTMLInputTypeAttribute,
    required?: boolean
    name?: string
    onChange?: React.ChangeEventHandler<HTMLInputElement>,
    value?: string
    error?: string | null
    autoComplete?: string
}

export function FormField(props: Props) {
  const { className, label, type, required, name, onChange, value, error, autoComplete } = props
  return  (
    <div className={classNames(s.Form__field, className)}>
      <label className={s.Form__label}>{label}</label>
      <input className={s.Form__input} type={type} required={required} name={name} onChange={onChange} autoComplete={autoComplete} value={value}/>
      <div className={s.error}>{error}</div>
    </div>
  )
}