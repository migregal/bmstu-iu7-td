import s from "./DashboardHomePage.module.css"
import Dropzone from "components/Dropzone"
import Placeholder from "./components/Placeholder"
import { DraftFilesContextProvider, useDraftFilesContext } from "./contexts/DraftFilesContext"
import DraftFilesTable from "./components/DraftFilesTable"

const ACCEPT = {"text/markdown": [".md"]}

function DashboardDropzone() {
  const { isLoading, handleDrop } = useDraftFilesContext()

  return <Dropzone accept={ACCEPT} onDrop={handleDrop} title="Drop 🗏 or 🗐 here" disabled={isLoading} />
}

export function DashboardHomePage() {
  return <DraftFilesContextProvider>
    <main className={s.root}>
      <DashboardDropzone />
      <Placeholder />
      <DraftFilesTable />
    </main>
  </DraftFilesContextProvider>
}
