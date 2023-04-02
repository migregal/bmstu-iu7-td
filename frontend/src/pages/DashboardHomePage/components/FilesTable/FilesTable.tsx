import { Table, TableBody, TableCell, TableRow } from "components/FilesTable"

import { useFilesContext } from "../../contexts/FilesContext"
import pluralizeFileSize from "../../helpers/pluralizeFileSize"

import styles from "./FilesTable.module.css"
import { ReactComponent as QRCode } from "./QRCode.svg"
import { ReactComponent as Delete } from "./Delete.svg"
import { ReactComponent as OpenFile } from "./OpenFile.svg"
import { ReactComponent as UpdateFile} from "./UpdateFile.svg"
import { ReactComponent as Edit} from "./Edit.svg"

export function FilesTable() {
  const { files } = useFilesContext()

  return <Table className={styles.table}>
    <TableBody>
      {files.map(file => (
        <TableRow key={file.id}>
          <TableCell className={styles.cell_qr}>
            <QRCode />
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
              <a href="#" target="_blank" className={styles.openFile}>
                <OpenFile />
              </a>
              <span className={styles.size}>{pluralizeFileSize(file.length)}</span>
            </div>
          </TableCell>
          <TableCell className={styles.cell_delete}>
            <Delete />
          </TableCell>
        </TableRow>
      ))}
    </TableBody>
  </Table>
}
