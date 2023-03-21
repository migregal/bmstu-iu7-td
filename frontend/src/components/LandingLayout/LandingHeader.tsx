import classnames from "classnames"
import { NavLink } from "react-router-dom"
import { PATH } from "routes/paths"
import s from "./LandingHeader.module.css"

type Props = {
    className?: string
}

export function LandingHeader({ className }: Props) {
  return <nav className={classnames(s.LandingHeader, className)}>
    <NavLink to={"#"} className={s.LandingHeader__link}>
        Sign in
    </NavLink>
    <NavLink to={PATH.REGISTRATION} className={s.LandingHeader__link}>
        Create account
    </NavLink>
  </nav>
}