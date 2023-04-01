import {  useCallback, useRef, useState, createContext, PropsWithChildren, useContext } from "react"
import { FileRejection } from "react-dropzone"

import bus from "bus"
import { FetchAddFileBody, fetchAddFile } from "api/fetchAddFile"
import { DraftFile } from "types/DraftFile"

import getFileBaseName from "../helpers/getFileBaseName"
import isValidDraft from "../helpers/isValidDraft"

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

export function createDraftFilesContext() {
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

    function fileToKey(file: File) {
      return file.name
    }

    const updatingFiles = new Set([
      ...acceptedFiles.map(fileToKey),
      ...fileRejections.map(({file}) => fileToKey(file))
    ])

    setDraftFiles((prev) => [
      ...fileRejections.map(({ file, errors }) => ({
        file,
        fileErrors: errors,
        title: getFileBaseName(file.name),
      })),
      ...acceptedFiles.map(file => ({
        file,
        title: getFileBaseName(file.name),
      })),
      ...prev.filter(draft => !updatingFiles.has(fileToKey(draft.file))),
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

    const created = new Map<DraftFile, FetchAddFileBody>()
    const failed = new Map<DraftFile, string[]>()

    for(const draft of draftFilesRef.current.filter(isValidDraft)) {
      const data = new FormData()
      data.append("title", draft.title)
      data.append("file", draft.file)

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

    if (created.size > 0) {
      bus.emit("filesCreated", {
        created: Array.from(created.entries())
          .map(([draft, data]) => ({ draft, data }))
      })
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

export function DraftFilesContextProvider({ children }: PropsWithChildren) {
  return <DraftFilesContext.Provider value={createDraftFilesContext()}>{children}</DraftFilesContext.Provider>
}
