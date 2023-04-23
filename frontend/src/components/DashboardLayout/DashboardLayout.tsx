import { Outlet } from "react-router-dom"
import s from "./DashboardLayout.module.css"
import { Header } from "./Header"

export function DashboardLayout() {
  return (<div className={s.DashboardLayout}>
    <Header className={s.DashboardLayout__header} /> 
    <div className={s.DashboardLayout__content}>
      <Outlet />
    </div>
  </div>)
}