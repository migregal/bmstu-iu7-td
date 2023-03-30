import classNames from "classnames"
import s from "./Placeholder.module.css"
import { ReactComponent as Empty } from "./Empty.svg"
import { useDraftFilesContext } from "pages/DashboardHomePage/contexts/DraftFilesContext"

type Props = {
    className?: string
}

export function Placeholder({ className }: Props) {
  const { draftFiles } = useDraftFilesContext()

  const shouldShow = draftFiles.length === 0

  return (
    <div className={classNames(s.placeholder__wrapper, className)}>
      <div className={classNames(s.placeholder, shouldShow && s.placeholder_show)}>
        <p className={s.placeholder__text}>You haven&apos;t created any files yet 🤷⬆️</p>
        <Empty className={s.placeholder__image}/>
      </div>
    </div>
  )
}
