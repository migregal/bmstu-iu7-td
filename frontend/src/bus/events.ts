import { FetchAddFileBody } from "api/fetchAddFile"
import { DraftFile } from "types/DraftFile"

export type FilesCreatedEventPayload = {
    created: { draft: DraftFile, data: FetchAddFileBody }[]
}

export type Events = {
  filesCreated: (payload: FilesCreatedEventPayload) => void
}
