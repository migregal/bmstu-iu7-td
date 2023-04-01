import Compose from "components/Compose"

import { DraftFilesContextProvider } from "./contexts/DraftFilesContext"
import { FilesContextProvider } from "./contexts/FilesContext"

import DraftFilesTable from "./components/DraftFilesTable"
import Dropzone from "./components/Dropzone"
import Placeholder from "./components/Placeholder"

import s from "./DashboardHomePage.module.css"
import FilesTable from "./components/FilesTable"

export function DashboardHomePage() {
  return <Compose
    components={[FilesContextProvider, DraftFilesContextProvider]}
  >
    <main className={s.root}>
      <Dropzone />
      <Placeholder />
      <DraftFilesTable />
      <FilesTable />
    </main>
  </Compose>
}
