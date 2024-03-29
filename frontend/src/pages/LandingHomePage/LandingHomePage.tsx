import { useViewerContext } from "contexts/viewer"
import { Link, Navigate } from "react-router-dom"
import { PATH } from "routes/paths"

export function LandingHomePage() {
  const { isAuth } = useViewerContext()

  if (isAuth) {
    return <Navigate to={PATH.DASHBOARD} replace />
  }

  return <main>
    <h1>Markup2</h1>
    <p>
      Some markdown pages for you ❤️
    </p>
    <p>
      <Link to={PATH.LOGIN}>Log in</Link> to your account if you already have one.
    </p>
    <p>
      Or <Link to={PATH.REGISTRATION}>create a new one</Link> if this is your first time here, sweetie ✨.
    </p>
  </main>
}
