import classNames from "classnames"

import Button from "components/Button"
import { Table, TableBody, TableCell, TableRow } from "components/FilesTable"

import { useDraftFilesContext } from "../../contexts/DraftFilesContext"
import isValidDraft from "../../helpers/isValidDraft"
import pluralizeFileSize from "../../helpers/pluralizeFileSize"

import s from "./DraftFilesTable.module.css"
import { ReactComponent as CloseIcon } from "./Close.svg"

export function DraftFilesTable() {
  const { draftFiles, handleChangeDraft, isLoading, handleDeleteDraft, handleSaveDraft, handleCancelDraft } = useDraftFilesContext()
  const canSaveDraft = draftFiles.find(isValidDraft)

  if (draftFiles.length === 0) {
    return null
  }

  return <>
    <Table className={s.table}>
      <TableBody>
        {draftFiles.map(({ file, title, fileErrors, saveErrors }, i) => (
          <TableRow className={s.row} key={i} data-index={i}>
            <TableCell className={s.cell_title}>
              <input value={title} className={s.title__input} onChange={handleChangeDraft}/>
              <ul className={s.title__errors}>
                {saveErrors?.map((error) => (
                  <li key={error}>
                    {error}
                  </li>
                ))}
                {fileErrors?.map((error) => (
                  <li key={error.code}>
                    {error.message}
                  </li>
                ))}
              </ul>
            </TableCell>
            <TableCell className={s.cell_size}>
              {pluralizeFileSize(file.size)}
            </TableCell>
            <TableCell className={s.cell_delete}>
              <CloseIcon
                className={classNames(s.deleteDraft, isLoading && s.deleteDraft_disabled)}
                onClick={handleDeleteDraft}
                title="Remove draft"
              />
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
    <div className={s.actions}>
      {canSaveDraft && <Button onClick={handleSaveDraft} disabled={isLoading}>Save</Button>}
      <Button outline onClick={handleCancelDraft}>Cancel</Button>
    </div>
  </>
}
