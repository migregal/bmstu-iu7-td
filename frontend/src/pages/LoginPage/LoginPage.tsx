import { fetchLogin } from "api/fetchLogin"
import Button from "components/Button"
import { Form, FormField } from "components/StraightForm"
import { useViewerContext } from "contexts/viewer"
import { useState } from "react"
import { Link, useNavigate } from "react-router-dom"
import { PATH } from "routes/paths"
import s from "./LoginPage.module.css"

export function LoginPage() {
  const [state, setState] = useState({
    login: "",
    password: "",
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
    const { login, password } = state
    setIsLoading(true)

    try {
      const { data, errors } = await fetchLogin({ login, password })

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
    } catch(error) {
      console.error("LoginPage.handleSubmit", error)
      setState(state => ({...state, errors: {default: "Unknown error" }}))
      setIsLoading(false)
    } 
  }

  return <main className={s.LoginPage}>
    <Form className={s.LoginPage__form} onSubmit={handleSubmit} title="Sign in" error={state.errors.default}>
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
      <Button disabled={isLoading}>Login</Button>
      <p>
        Or <Link to={PATH.REGISTRATION}>create new account</Link>
      </p>
    </Form>
  </main>
}
