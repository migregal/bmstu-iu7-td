import { Table, TableBody } from "components/FilesTable"

import { useFilesContext } from "../../contexts/FilesContext"

import styles from "./FilesTable.module.css"
import { FileRow } from "./FileRow"

export function FilesTable() {
  const { files, handlePartialChange, handleUploadFile, handleDeleteFile } = useFilesContext()

  return <Table className={styles.table}>
    <TableBody>
      {files.map(file => (
        <FileRow
          key={file.id}
          file={file}
          onChange={handlePartialChange}
          onChangeFile={handleUploadFile}
          onDelete={handleDeleteFile}
        />
      ))}
    </TableBody>
  </Table>
}
