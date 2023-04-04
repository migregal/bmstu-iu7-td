import React, { useCallback } from "react"

import { TableCell, TableRow } from "components/FilesTable"

import { File } from "../../contexts/FilesContext"
import pluralizeFileSize from "../../helpers/pluralizeFileSize"

import EditableTitle from "../EditableTitle"
import QRCodeButton from "../QRCodeButton"

import styles from "./FilesTable.module.css"
import { ReactComponent as Delete } from "./Delete.svg"
import { ReactComponent as OpenFile } from "./OpenFile.svg"
import { ReactComponent as UpdateFile} from "./UpdateFile.svg"

type Props = {
    file: File
    onChange?: (id: string, update: Partial<File>) => Promise<{ errors?: Record<string, string> | null }>
}

export const FileRow = React.memo(function FileRow(props: Props) {
  const { file, onChange } = props

  const handleSaveTitle = useCallback(async (title: string) => {
    if (onChange) {
      return onChange(file.id, { title })
    }
    return {}
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
        <div className={styles.updateFile}>
          <input type="file" accept="text/markdown,.md" />
          <UpdateFile />
        </div>
        <a href={`${file.url}?format=md`} target="_blank" className={styles.openFile} rel="noreferrer">
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
