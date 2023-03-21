import DashboardLayout from "components/DashboardLayout"
import RequireViewer from "components/RequireViewer"
import DashboardHomePage from "pages/DashboardHomePage"
import RegistrationPage from "pages/RegistrationPage"
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
      }
    ]
  },
  {
    path: PATH.DASHBOARD,
    element: <DashboardLayout />,
    children: [
      {
        index: true,
        element: <RequireViewer><DashboardHomePage /></RequireViewer>,
      }
    ]
  }
])