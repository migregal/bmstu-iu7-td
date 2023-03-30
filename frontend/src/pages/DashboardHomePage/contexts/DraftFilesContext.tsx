import {  useCallback, useRef, useState, createContext, PropsWithChildren, useContext } from "react"
import { FileError, FileRejection } from "react-dropzone"
import { fetchAddFile } from "api/fetchAddFile"

export type DraftFile = {
  file: File
  title: string
  fileErrors?: FileError[]
  saveErrors?: string[] | null
}

type Params = {
    onCreate?: () => void;
}

export function getFileBaseName(name: string) {
  return name.split(/\.(?=[^.]+$)/)[0]
}

export function isValidDraft(draft: DraftFile) {
  return !draft.fileErrors || !draft.fileErrors.length
}

function errorsObjectToArray(errors: Record<string, string>, knownFields: string[] = [], onlyKnown = false) {
  errors = {...errors}
  const arr = []

  for (const key of knownFields) {
    if (key in errors) {
      arr.push(errors[key])
      delete errors[key]
    }
  }

  return onlyKnown ? arr : arr.concat(Object.values(errors))
}

export function createDraftFilesContext({ onCreate }: Params) {
  const [draftFiles, setDraftFiles] = useState<DraftFile[]>([])
  const draftFilesRef = useRef<DraftFile[]>(draftFiles)
  draftFilesRef.current = draftFiles

  const [isLoading, setIsLoading] = useState(false)

  function getDraftFileFromElement(target: Element) {
    const i = target.closest("[data-index]")?.getAttribute("data-index")
    if (!i) {
      return null
    }

    return draftFilesRef.current[+i]
  }

  const handleDrop = useCallback((
    acceptedFiles: File[],
    fileRejections: FileRejection[],
  ) => {
    setDraftFiles((prev) => [
      ...fileRejections.map(({ file, errors }) => ({
        file,
        errors,
        title: getFileBaseName(file.name),
      })),
      ...acceptedFiles.map(file => ({
        file,
        title: getFileBaseName(file.name),
      })),
      ...prev,
    ])
  }, [])

  const handleChangeDraft: React.ChangeEventHandler<HTMLInputElement> = useCallback((event) => {
    const target = event.target
    const draft = getDraftFileFromElement(target)
    const value = target.value

    if (!draft) {
      return
    }

    setDraftFiles((drafts) => drafts.map(d => d === draft ? { ...d, title: value } : d ))
  }, [])

  const handleDeleteDraft: React.MouseEventHandler<SVGSVGElement> = useCallback((event) => {
    const draft = getDraftFileFromElement(event.target as SVGSVGElement)
    if (!draft) {
      return
    }

    setDraftFiles((drafts) => drafts.filter(d => d !== draft))
  }, [])

  const handleSaveDraft = useCallback(async () => {
    setIsLoading(true)

    const created = new Map<DraftFile, { id: string, url: string }>()
    const failed = new Map<DraftFile, string[]>()

    for(const draft of draftFilesRef.current.filter(isValidDraft)) {
      const data = new FormData()
      data.append("title", draft.title)
      data.append("file", draft.file)

      await new Promise((res) => setTimeout(res, 5000))

      try {
        const result = await fetchAddFile(data)

        if (result.data) {
          created.set(
            draft,
            result.data,
          )
        } else {
          failed.set(
            draft,
            errorsObjectToArray(result.errors || { default: "Wrong response"}, ["default", "title", "file"]),
          )
        }
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      } catch (error: any) {
        failed.set(draft, [error.message || "Unknown error"])
      }
    }

    if (onCreate) {
      onCreate() // TODO: pass data created
    }

    setDraftFiles(drafts =>
      drafts
        .filter(draft => !created.has(draft))
        .map(draft => failed.has(draft) ? {...draft, saveErrors: failed.get(draft) } : draft)
    )
    setIsLoading(false)
  }, [])

  const handleCancelDraft = useCallback(() => {
    setDraftFiles([])
  }, [])

  return {
    draftFiles,
    setDraftFiles,
    isLoading,
    handleDrop,
    handleChangeDraft,
    handleSaveDraft,
    handleDeleteDraft,
    handleCancelDraft,
  }
}

export type DraftFilesContextValue = ReturnType<typeof createDraftFilesContext>

export const DraftFilesContext = createContext<DraftFilesContextValue | null>(null)

export function useDraftFilesContext() {
  return useContext(DraftFilesContext)!
}

export function DraftFilesContextProvider({ children, ...rest }: PropsWithChildren<Params>) {
  return <DraftFilesContext.Provider value={createDraftFilesContext(rest)}>{children}</DraftFilesContext.Provider>
}
