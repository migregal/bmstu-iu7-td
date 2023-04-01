import InnerDropzone from "components/Dropzone"
import { useDraftFilesContext } from "../contexts/DraftFilesContext"

const ACCEPT = {"text/markdown": [".md"]}

export function Dropzone() {
  const { isLoading, handleDrop } = useDraftFilesContext()

  return <InnerDropzone
    accept={ACCEPT}
    onDrop={handleDrop}
    title="Drop 🗏 or 🗐 here"
    disabled={isLoading}
  />
}

export default Dropzone

