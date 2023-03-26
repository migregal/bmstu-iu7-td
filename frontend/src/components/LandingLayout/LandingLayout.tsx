import { Outlet } from "react-router-dom"
import s from "./LandingLayout.module.css"

export function LandingLayout() {
  return (
    <div className={s.LandingLayout}>
      <div className={s.LandingLayout__content}>
        <Outlet />
      </div>
    </div>
  )
}