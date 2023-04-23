import { fetchAllFiles } from "api/fetchAllFiles"
import { fetchDeleteFile } from "api/fetchDeleteFile"
import { fetchUpdateFile } from "api/fetchUpdateFile"
import bus from "bus"
import { DraftsSavedEventPayload } from "bus/events"
import { useViewerContext } from "contexts/viewer"
import { useRef, useState, createContext, PropsWithChildren, useContext, useEffect, useCallback } from "react"

export type File = {
  id: string
  url: string
  length: number
  title: string
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
          if (errors.default === "unauthorized") {
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
    function onFilesCreated({ saved: created }: DraftsSavedEventPayload) {
      setFiles(files => [
        ...created.map(({ draft, data }) => ({
          ...data,
          title: draft.title,
          length: draft.file.size,
        })),
        ...files,
      ])
    }

    bus.addListener("draftsSaved", onFilesCreated)

    return () => {
      bus.removeListener("draftsSaved", onFilesCreated)
    }
  }, [])

  const handlePartialChange = useCallback(async (id: string, update: Partial<File>) => {
    const form = new FormData()

    if (update.title) {
      form.append("title", update.title)
    }

    try {
      const { data, errors } = await fetchUpdateFile(id, form)
      if (data) {
        setFiles(files => files.map(file => file.id === id ? { ...file, ...update } : file))
        return {}
      } else if (errors) {
        return { errors }
      } else {
        console.error("createFilesContext.handlePartialChange unknown response")
        return { errors: {default: "Unknown repsonse"}}
      }
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } catch (error: any) {
      console.error("createFilesContext.handlePartialChange", error)
      return { errors: {default: error.message || "Unknown error"}}
    }
  }, [])

  const handleUploadFile = useCallback(async (id: string, file: globalThis.File) => {
    const form = new FormData()
    form.append("file", file)

    try {
      const { data, errors } = await fetchUpdateFile(id, form)
      if (data) {
        setFiles(files => files.map(f => f.id === id ? { ...f, length: file.size } : f))
        return {}
      } else if (errors) {
        return { errors }
      } else {
        console.error("createFilesContext.handleUploadFile unknown response")
        return { errors: {default: "Unknown repsonse"}}
      }
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } catch (error: any) {
      console.error("createFilesContext.handleUploadFile", error)
      return { errors: {default: error.message || "Unknown error"}}
    }
  }, [])

  const handleDeleteFile = useCallback(async (id: string) => {
    try {
      const { data, errors } = await fetchDeleteFile(id)
      if (data) {
        setFiles(files => files.filter(file => file.id !== id))
        return {}
      } else if (errors) {
        return { errors }
      } else {
        console.error("createFilesContext.deleteFile unknown response")
        return { errors: { default: "Unknown response" }}
      }
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } catch (error: any) {
      console.error("createFilesContext.deleteFile", error)
      return { errors: {default: error.message || "Unknown error"}}
    }
  }, [])

  return {
    files,
    loadErrors,
    isLoading,
    handleUploadFile,
    handlePartialChange,
    handleDeleteFile,
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
