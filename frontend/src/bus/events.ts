import { FetchAddFileBody } from "api/fetchAddFile"
import { DraftFile } from "types/DraftFile"

export type DraftsSavedEventPayload = {
    saved: { draft: DraftFile, data: FetchAddFileBody }[]
}

export type Events = {
  draftsSaved: (payload: DraftsSavedEventPayload) => void
}
