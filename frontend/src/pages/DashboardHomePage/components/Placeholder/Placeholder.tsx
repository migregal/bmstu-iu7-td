import InnerPlaceholder from "components/Placeholder"

import { useDraftFilesContext } from "../../contexts/DraftFilesContext"
import { useFilesContext } from "../../contexts/FilesContext"

import { ReactComponent as Empty } from "./Empty.svg"

type Props = {
    className?: string
}

export function Placeholder({ className }: Props) {
  const { draftFiles } = useDraftFilesContext()
  const { files, isLoading } = useFilesContext()

  return <InnerPlaceholder
    shouldShow={draftFiles.length === 0 && files.length === 0 && !isLoading}
    className={className}
    message="You haven't created any files yet 🤷⬆️"
  >
    <Empty />
  </InnerPlaceholder>
}
