import DashboardLayout from "components/DashboardLayout"
import RequireViewer from "components/RequireViewer"
import DashboardHomePage from "pages/DashboardHomePage"
import LoginPage from "pages/LoginPage"
import RegistrationPage from "pages/RegistrationPage"
import NotFoundPage from "pages/NotFoundPage"
import {
  createBrowserRouter,
} from "react-router-dom"
import LandingLayout from "../components/LandingLayout"
import LandingHomePage from "../pages/LandingHomePage"
import { PATH } from "./paths"

export const router = createBrowserRouter([
  {
    path: PATH.INDEX,
    element: <LandingLayout />,
    children: [
      {
        index: true,
        element: <LandingHomePage />,
      },
      {
        path: PATH.REGISTRATION,
        element: <RegistrationPage />
      },
      {
        path: PATH.LOGIN,
        element: <LoginPage />
      },
      {
        path: PATH.NOT_FOUND,
        element: <NotFoundPage />
      }
    ]
  },
  {
    path: PATH.DASHBOARD,
    element: <RequireViewer><DashboardLayout /></RequireViewer>,
    children: [
      {
        index: true,
        element: <DashboardHomePage />,
      }
    ]
  },
])
