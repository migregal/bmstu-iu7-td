import { ReactComponent as Empty } from "./Empty.svg"
import s from "./DashboardHomePage.module.css"
import Dropzone from "components/Dropzone"

const ACCEPT = {"text/markdown": [".md"]}

export function DashboardHomePage() {
  return <main className={s.root}>
    <Dropzone accept={ACCEPT}/>
    <div className={s.placeholder}>
      <p className={s.placeholder__text}>You haven&apos;t created any files yet 🤷</p>
      <Empty className={s.placeholder__image}/>
    </div>
  </main>
}