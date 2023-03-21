import { ViewerContextProvider } from "contexts/viewer"
import { RouterProvider } from "react-router-dom"
import { router } from "./routes/router"

function App() {  
  return (
    <ViewerContextProvider>
      <RouterProvider router={router} />
    </ViewerContextProvider>
  )
}

export default App
