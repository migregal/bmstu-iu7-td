import classNames from "classnames"
import { PropsWithChildren } from "react"
import styles from "./FilesTable.module.css"

type Props = {
    className?: string
}

export function Table(props: PropsWithChildren<Props>) {
  const { children, className } = props
  return <table className={classNames(styles.table, className)}>{children}</table>
}
