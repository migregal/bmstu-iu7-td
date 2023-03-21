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
      }
    ]
  }
])