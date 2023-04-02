import classNames from "classnames"
import styles from "./FilesTable.module.css"

export function TableRow(props: React.HTMLAttributes<HTMLTableRowElement>) {
  const { children, className, ...rest } = props
  return <tr className={classNames(styles.row, className)} {...rest}>{children}</tr>
}
