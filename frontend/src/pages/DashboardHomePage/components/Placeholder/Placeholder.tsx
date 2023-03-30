import classNames from "classnames"
import s from "./Placeholder.module.css"
import { ReactComponent as Empty } from "./Empty.svg"
import { useDraftFilesContext } from "pages/DashboardHomePage/contexts/DraftFilesContext"
import React from "react"

type Props = {
    className?: string
}

type InnerProps = {
  shouldShow: boolean
} & Props

const InnerPlaceholder = React.memo(function InnerPlaceholder({ className, shouldShow }: InnerProps) {
  return (
    <div className={classNames(s.placeholder__wrapper, className)}>
      <div className={classNames(s.placeholder, shouldShow && s.placeholder_show)}>
        <p className={s.placeholder__text}>You haven&apos;t created any files yet 🤷⬆️</p>
        <Empty className={s.placeholder__image}/>
      </div>
    </div>
  )
})

export function Placeholder(props: Props) {
  const { draftFiles } = useDraftFilesContext()

  return <InnerPlaceholder shouldShow={draftFiles.length === 0} {...props} />
}
