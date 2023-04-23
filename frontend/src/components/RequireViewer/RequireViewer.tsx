import { useViewerContext } from "contexts/viewer"
import { PropsWithChildren } from "react"
import { Navigate, useLocation } from "react-router-dom"


export function RequireViewer({ children }: PropsWithChildren) {
  const { isAuth } = useViewerContext()
  const location = useLocation()

  if (!isAuth) {
    return <Navigate to={"/"} state={{ from: location }} replace />
  }

  return <>{children}</>
}