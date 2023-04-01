import { DraftFile } from "types/DraftFile"

export function isValidDraft(draft: DraftFile) {
  return !draft.fileErrors || !draft.fileErrors.length
}

export default isValidDraft
