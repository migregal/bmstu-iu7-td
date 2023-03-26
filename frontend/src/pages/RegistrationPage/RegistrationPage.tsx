import { fetchRegistration } from "api/fetchRegistration"
import classNames from "classnames"
import Button from "components/Button"
import { useViewerContext } from "contexts/viewer"
import { useState } from "react"
import { Link, useNavigate } from "react-router-dom"
import { PATH } from "routes/paths"
import s from "./RegistrationPage.module.css"

export function RegistrationPage() {
  const [state, setState] = useState({
    login: "",
    password: "",
    confirmPassword: "",
    errors: {} as Record<string,  string | null>
  })
  const [isLoading, setIsLoading] = useState(false)
  const navigate = useNavigate()
  const { setViewer } = useViewerContext()

  const handleChange: React.ChangeEventHandler<HTMLInputElement>  = (event) => {
    const { name, value } = event.target
    setState(state => ({ ...state, [name]: value, errors: {...state.errors, [name]: null}}))
  }

  const handleSubmit: React.FormEventHandler<HTMLFormElement> = async (event) => {
    event.preventDefault()
    event.stopPropagation()
    const { login, password, confirmPassword } = state
    if (password !== confirmPassword) {
      setState(state => ({...state, errors: {...state.errors, confirmPassword: "It is not equal to password"}}))
      return
    }
    setIsLoading(true)

    try {
      const { data, errors } = await fetchRegistration({ login, password })

      if (data) {
        setViewer(data.id, data.token)
        navigate(PATH.DASHBOARD)
      }
      else if (errors) {
        setState(state => ({ ...state, errors }))
        setIsLoading(false)
      }
      else {
        setState(state => ({...state, errors: {default: "Unknown response" }}))
        setIsLoading(false)
      }
    } catch(error: any) {
      console.error("RegistrationPage.handleSubmit", error)
      setState(state => ({...state, errors: {default: "Unknown error" }}))
      setIsLoading(false)
    } 
  }

  return <main className={s.RegistrationPage}>
    <form className={classNames(s.Form, s.RegistrationPage__form)} onSubmit={handleSubmit}>
      <h2 className={s.Form}>Create account</h2>
      {state.errors.default && <p className={classNames(s.Form__group, s.error)}>{state.errors.default}</p>}
      <div className={s.Form__group}>
        <label className={s.Form__label}>Login</label>
        <input className={s.Form__input} required name="login" onChange={handleChange} value={state.login}/>
        <div className={s.error}>{state.errors.login}</div>
      </div>
      <div className={s.Form__group}>
        <label className={s.Form__label}>Password</label>
        <input className={s.Form__input} required type="password" name="password" onChange={handleChange} value={state.password}/>
        <div className={s.error}>{state.errors.password}</div>
      </div>
      <div className={s.Form__group}>
        <label className={s.Form__label}>Confirm password</label>
        <input className={s.Form__input} required type="password" autoComplete="false" name="confirmPassword" onChange={handleChange} value={state.confirmPassword}/>
        <div className={s.error}>{state.errors.confirmPassword}</div>
      </div>
      <Button className={s.Form__action} disabled={isLoading}>Create new account</Button>
      <div>
        Or <Link to={PATH.SIGN_IN}>log in to existing account</Link>
      </div>
    </form>
  </main>
}
