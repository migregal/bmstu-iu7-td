import { fetchAllFiles } from "api/fetchAllFiles"
import bus from "bus"
import { FilesCreatedEventPayload } from "bus/events"
import { useViewerContext } from "contexts/viewer"
import { useRef, useState, createContext, PropsWithChildren, useContext, useEffect } from "react"

export type File = {
  id: string
  url: string
  length: number
  title: string
  saveErrors?: string[] | null
}

export function createFilesContext() {
  const [files, setFiles] = useState<File[]>([])
  const [loadErrors, setLoadErrors] = useState<Record<string, string> | null>(null)
  const filesRef = useRef<File[]>(files)
  const { resetViewer } = useViewerContext()
  filesRef.current = files

  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    fetchAllFiles().then(
      ({ data, errors }) => {
        if (data) {
          setFiles(data.files.reverse())
          setIsLoading(false)
          setLoadErrors(null)
        }
        else if (errors) {
          setIsLoading(false)
          if (errors.id === "empty") {
            resetViewer()
          } else {
            setLoadErrors(errors)
          }
        } else {
          setIsLoading(false)
          setLoadErrors({default: "Unknown response. Try again later"})
        }
      },
      error => {
        console.error("createFilesContext.fetchAllFiles", error)
        setIsLoading(false)
        setLoadErrors({default: error.message || "Unknown error"})
      }
    )
  }, [])

  useEffect(() => {
    function onFilesCreated({ created }: FilesCreatedEventPayload) {
      setFiles(files => [
        ...created.map(({ draft, data }) => ({
          ...data,
          title: draft.title,
          length: draft.file.size,
        })),
        ...files,
      ])
    }

    bus.addListener("filesCreated", onFilesCreated)

    return () => {
      bus.removeListener("filesCreated", onFilesCreated)
    }
  }, [])

  return {
    files,
    loadErrors,
    isLoading,
  }
}

export type FilesContextValue = ReturnType<typeof createFilesContext>

export const FilesContext = createContext<FilesContextValue | null>(null)

export function useFilesContext() {
  return useContext(FilesContext)!
}

export function FilesContextProvider({ children }: PropsWithChildren) {
  return <FilesContext.Provider value={createFilesContext()}>{children}</FilesContext.Provider>
}
