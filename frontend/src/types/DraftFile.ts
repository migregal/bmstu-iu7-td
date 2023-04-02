import { FileError } from "react-dropzone"

export type DraftFile = {
  file: File
  title: string
  fileErrors?: FileError[]
  saveErrors?: string[] | null
}
