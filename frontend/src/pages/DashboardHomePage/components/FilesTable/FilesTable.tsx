import { useFilesContext } from "pages/DashboardHomePage/contexts/FilesContext"

export function FilesTable() {
  const { files } = useFilesContext()

  return <pre>
    {JSON.stringify(files, null, 2)}
  </pre>
}
