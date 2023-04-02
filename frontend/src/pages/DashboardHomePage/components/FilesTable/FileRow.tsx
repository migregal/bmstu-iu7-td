import React from "react"

import { TableCell, TableRow } from "components/FilesTable"

import { File } from "../../contexts/FilesContext"
import pluralizeFileSize from "../../helpers/pluralizeFileSize"

import QRCodeButton from "../QRCodeButton"

import styles from "./FilesTable.module.css"
import { ReactComponent as Delete } from "./Delete.svg"
import { ReactComponent as OpenFile } from "./OpenFile.svg"
import { ReactComponent as UpdateFile} from "./UpdateFile.svg"
import { ReactComponent as Edit} from "./Edit.svg"

type Props = {
    file: File
}

export const FileRow = React.memo(function FileRow(props: Props) {
  const { file } = props
  return <TableRow key={file.id}>
    <TableCell className={styles.cell_qr}>
      <QRCodeButton url={file.url} />
    </TableCell>
    <TableCell>
      <div className={styles.title}>
        <a href={file.url} target="_blank" rel="noreferrer">{file.title}</a>
        <Edit className={styles.edit}/>
      </div>
    </TableCell>
    <TableCell className={styles.cell_size}>
      <div className={styles.size__wrapper}>
        <div className={styles.updateFile}>
          <input type="file" accept="text/markdown,.md"/>
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
