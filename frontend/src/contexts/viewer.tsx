import { client } from "api/client"
import { createContext, PropsWithChildren, useContext, useState } from "react"

function createViewerContext() {
  const [viewer, _setViewer] = useState<{ id: string } | null>()

  const setViewer = (id: string, token: string) => {
    client.setToken(token)
    _setViewer({ id })
  }

  const resetViewer = () => {
    client.setToken(null)
    _setViewer(null)
  }

  return {
    isAuth: !!client.token,
    viewer,
    setViewer,
    resetViewer,
  }
}

export const viewerContext = createContext<ReturnType<typeof createViewerContext> | null>(null)

export function useViewerContext() {
  return useContext(viewerContext)!
}

export function ViewerContextProvider({ children }: PropsWithChildren) {
  return <viewerContext.Provider value={createViewerContext()}>{children}</viewerContext.Provider>
}
