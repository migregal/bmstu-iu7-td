import { fetchLogout } from "api/fetchLogout"
import classnames from "classnames"
import Button from "components/Button"
import { useViewerContext } from "contexts/viewer"
import { useNavigate } from "react-router-dom"
import { PATH } from "routes/paths"
import s from "./Header.module.css"

type Props = {
    className?: string
}

export function Header({ className }: Props) {

  const { resetViewer } = useViewerContext()
  const navigate = useNavigate()

  const handleLogout = () => {
    resetViewer()
    fetchLogout()
    navigate(PATH.INDEX)
  }

  return <nav className={classnames(s.Header, className)}>
    <Button className={s.Header__logout} onClick={handleLogout}>
        Log out
    </Button >
  </nav>
}