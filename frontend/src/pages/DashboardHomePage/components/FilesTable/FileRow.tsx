import React, { useCallback } from "react"
import { toast } from "react-toastify"

import { TableCell, TableRow } from "components/FilesTable"

import { File } from "../../contexts/FilesContext"
import pluralizeFileSize from "../../helpers/pluralizeFileSize"
import errorsObjectToArray from "../../helpers/errorsObjectToArray"

import EditableTitle from "../EditableTitle"
import QRCodeButton from "../QRCodeButton"
import UpdateFileButton from "../UpdateFileButton"
import DeleteForeverButton from "../DeleteForeverButton"

import styles from "./FilesTable.module.css"
import { ReactComponent as OpenFile } from "./OpenFile.svg"
import IconButton from "components/IconButton"

type Props = {
    file: File
    onChange?: (id: string, update: Partial<File>) => Promise<{ errors?: Record<string, string> | null }>
    onChangeFile?: (id: string, file: globalThis.File) => Promise<{ errors?: Record<string, string> | null }>
    onDelete?: (id: string) => Promise<{ errors?: Record<string, string> | null }>
}

export const FileRow = React.memo(function FileRow(props: Props) {
  const { file, onChange, onChangeFile, onDelete } = props

  const handleSaveTitle = useCallback(async (title: string) => {
    if (onChange) {
      return onChange(file.id, { title })
    }
    return {}
  }, [file.id])

  const handleDrop = useCallback(async (newFile: globalThis.File) => {
    if (!onChangeFile) {
      return {}
    }
    return onChangeFile(file.id, newFile)
  }, [file.id])

  const handleDeleteFile = useCallback(async () => {
    if (!onDelete) {
      return {}
    }
    const { errors } = await onDelete(file.id)

    if (errors) {
      toast.error(errorsObjectToArray(errors, ["default"]).join(". "))
    } else {
      toast.success("Deleted!")
    }
  }, [file.id])

  return <TableRow key={file.id} data-id={file.id}>
    <TableCell className={styles.cell_qr}>
      <QRCodeButton url={file.url} />
    </TableCell>
    <TableCell>
      <EditableTitle
        value={file.title}
        href={file.url}
        onSave={handleSaveTitle}
      />
    </TableCell>
    <TableCell className={styles.cell_size}>
      <div className={styles.size__wrapper}>
        <UpdateFileButton onDrop={handleDrop} />
        <IconButton renderAs="a" href={`${file.url}?format=md`} target="_blank" className={styles.openFile} rel="noreferrer" title="Open source file">
          <OpenFile />
        </IconButton>
        <span className={styles.size}>{pluralizeFileSize(file.length)}</span>
      </div>
    </TableCell>
    <TableCell className={styles.cell_delete}>
      <div className={styles.delete__wrapper}>
        <div className={styles.delete}>
          <DeleteForeverButton onClick={handleDeleteFile}/>
        </div>
      </div>
    </TableCell>
  </TableRow>
})
