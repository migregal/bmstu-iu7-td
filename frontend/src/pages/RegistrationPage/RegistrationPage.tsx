import { fetchRegistration } from "api/fetchRegistration"
import Button from "components/Button"
import { Form, FormField } from "components/StraightForm"
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

  const handleSubmit: React.FormEventHandler<HTMLFormElement> = async () => {
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
        setTimeout(() => navigate(PATH.DASHBOARD), 0)
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
    <Form className={s.RegistrationPage__form} onSubmit={handleSubmit} title="Create account" error={state.errors.default}>
      <FormField 
        label="Email"
        type="email"
        required
        onChange={handleChange}
        name="login"
        value={state.login}
        error={state.errors.login}
      />
      <FormField 
        label="Password"
        type="password"
        required
        onChange={handleChange}
        name="password"
        value={state.password}
        error={state.errors.password}
      />
      <FormField 
        label="Confirm password"
        type="password"
        required
        autoComplete="false"
        onChange={handleChange}
        name="confirmPassword"
        value={state.confirmPassword}
        error={state.errors.confirmPassword}
      />
      <Button disabled={isLoading}>Create new account</Button>
      <p>
        Or <Link to={PATH.LOGIN}>log in to existing account</Link>
      </p>
    </Form>
  </main>
}
