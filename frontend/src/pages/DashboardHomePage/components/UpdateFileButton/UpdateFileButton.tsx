import React, { useState } from "react"
import { useDropzone } from "react-dropzone"
import { toast } from "react-toastify"
import classNames from "classnames"

import IconButton from "components/IconButton"
import { ACCEPT } from "constants/accepts"

import errorsObjectToArray from "../../helpers/errorsObjectToArray"

import styles from "./UpdateFileButton.module.css"
import { ReactComponent as UpdateFile} from "./UpdateFile.svg"

type Props = {
    className?: string
    onDrop?: (file: File) => Promise<{ errors?: Record<string, string> | null }>
}

export const UpdateFileButton = React.memo(function UpdateFileButton (props: Props) {
  const { className, onDrop } = props

  const [loading, setLoading] = useState(false)

  const { getInputProps, getRootProps, } = useDropzone({
    accept: ACCEPT,
    maxFiles: 1,
    async onDrop(acceptedFiles, fileRejections) {
      if (fileRejections.length > 0) {
        toast.error(fileRejections.flatMap(({ errors }) => errors).map(error => error.message).join(". "))
      }

      if (!acceptedFiles.length || !onDrop) {
        return{}
      }

      setLoading(true)

      try {
        const { errors } = await onDrop(acceptedFiles[0])

        if (errors) {
          toast.error(errorsObjectToArray(errors, ["default", "file"]).join(". "))
        } else {
          toast.success("Saved!")
        }

      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      } catch (error: any) {
        console.error("UpdateFileButton.onDrop", error)
        toast.error(error.message || "Unknown error")
      } finally {
        setLoading(false)
      }
    },
  })

  return <IconButton {...getRootProps({ className: classNames(styles.updateFile, className), disabled: loading })}>
    <input {...getInputProps()}/>
    <UpdateFile />
  </IconButton>
})
