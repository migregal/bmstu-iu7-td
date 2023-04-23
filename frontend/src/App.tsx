import { ToastContainer } from "react-toastify"
import { RouterProvider } from "react-router-dom"

import { ViewerContextProvider } from "contexts/viewer"
import { router } from "./routes/router"

import "react-toastify/dist/ReactToastify.css"

function App() {
  return (
    <ViewerContextProvider>
      <RouterProvider router={router} />
      <ToastContainer />
    </ViewerContextProvider>
  )
}

export default App
