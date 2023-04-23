import classNames from "classnames"
import { PropsWithChildren } from "react"
import styles from "./FilesTable.module.css"

type Props = {
    className?: string
}

export function TableCell(props: PropsWithChildren<Props>) {
  const { children, className } = props
  return <td className={classNames(styles.cell, className)}>{children}</td>
}
