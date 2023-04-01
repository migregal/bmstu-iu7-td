import classNames from "classnames"

import Button from "components/Button"

import { useDraftFilesContext } from "../../contexts/DraftFilesContext"
import isValidDraft from "../../helpers/isValidDraft"

import s from "./DraftFilesTable.module.css"
import { ReactComponent as CloseIcon } from "./Close.svg"

const FILE_SIZES = [
  "b",
  "kb",
  "mb",
  "gb",
]

function pluralizeFileSize(size: number) {
  let i = 0
  while (i < FILE_SIZES.length - 1 && Math.ceil(size) > 1024) {
    size /= 1024
    i++
  }
  return Math.ceil(size) + FILE_SIZES[i]
}

export function DraftFilesTable() {
  const { draftFiles, handleChangeDraft, isLoading, handleDeleteDraft, handleSaveDraft, handleCancelDraft } = useDraftFilesContext()
  const canSaveDraft = draftFiles.find(isValidDraft)

  if (draftFiles.length === 0) {
    return null
  }

  return <>
    <table className={s.table}>
      <tbody>
        {draftFiles.map(({ file, title, fileErrors, saveErrors }, i) => (
          <tr className={s.row} key={i} data-index={i}>
            <td className={classNames(s.cell, s.cell_qr)}></td>
            <td className={classNames(s.cell, s.cell_title)}>
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
            </td>
            <td className={classNames(s.cell, s.cell_size)}>
              {pluralizeFileSize(file.size)}
            </td>
            <td className={classNames(s.cell, s.cell_delete)}>
              <CloseIcon
                className={classNames(s.deleteDraft, isLoading && s.deleteDraft_disabled)}
                onClick={handleDeleteDraft}
                title="Remove draft"
              />
            </td>
          </tr>
        ))}
      </tbody>
    </table>
    <div className={s.actions}>
      {canSaveDraft && <Button onClick={handleSaveDraft} disabled={isLoading}>Save</Button>}
      <Button outline onClick={handleCancelDraft}>Cancel</Button>
    </div>
  </>
}
