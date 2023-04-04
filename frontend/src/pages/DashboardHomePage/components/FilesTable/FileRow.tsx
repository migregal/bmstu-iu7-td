import React, { useCallback } from "react"

import { TableCell, TableRow } from "components/FilesTable"

import { File } from "../../contexts/FilesContext"
import pluralizeFileSize from "../../helpers/pluralizeFileSize"

import EditableTitle from "../EditableTitle"
import QRCodeButton from "../QRCodeButton"
import UpdateFileButton from "../UpdateFileButton"

import styles from "./FilesTable.module.css"
import { ReactComponent as Delete } from "./Delete.svg"
import { ReactComponent as OpenFile } from "./OpenFile.svg"

type Props = {
    file: File
    onChange?: (id: string, update: Partial<File>) => Promise<{ errors?: Record<string, string> | null }>
    onChangeFile?: (id: string, file: globalThis.File) => Promise<{ errors?: Record<string, string> | null }>
}

export const FileRow = React.memo(function FileRow(props: Props) {
  const { file, onChange, onChangeFile } = props

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
        <a href={`${file.url}?format=md`} target="_blank" className={styles.openFile} rel="noreferrer" title="Open source file">
          <OpenFile />
        </a>
        <span className={styles.size}>{pluralizeFileSize(file.length)}</span>
      </div>
    </TableCell>
    <TableCell className={styles.cell_delete}>
      <Delete />
    </TableCell>
  </TableRow>
})
