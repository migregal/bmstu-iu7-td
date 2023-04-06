import InnerDropzone from "components/Dropzone"
import { ACCEPT } from "constants/accepts"

import { useDraftFilesContext } from "../contexts/DraftFilesContext"

export function Dropzone() {
  const { isLoading, handleDrop } = useDraftFilesContext()

  return <InnerDropzone
    accept={ACCEPT}
    onDrop={handleDrop}
    title="Drop files here"
    disabled={isLoading}
  />
}

export default Dropzone

