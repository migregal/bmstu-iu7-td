import { Table, TableBody } from "components/FilesTable"

import { useFilesContext } from "../../contexts/FilesContext"

import styles from "./FilesTable.module.css"
import { FileRow } from "./FileRow"

export function FilesTable() {
  const { files } = useFilesContext()

  return <Table className={styles.table}>
    <TableBody>
      {files.map(file => (
        <FileRow file={file} key={file.id} />
      ))}
    </TableBody>
  </Table>
}
