import React, { useRef, useState } from "react"
import classNames from "classnames"

import IconButton from "components/IconButton"

import styles from "./DeleteForeverButton.module.css"
import { ReactComponent as Delete } from "./Delete.svg"
import { ReactComponent as DeleteForever } from "./DeleteForever.svg"

type Props = {
    className?: string
    onClick?: () => void;
}

export const DeleteForeverButton = React.memo(function DeleteForeverButton(props: Props) {
  const [open, setOpen] = useState<boolean>(false)
  const timeoutRef = useRef<number | null>(null)

  const handleClick = () => {
    if (!open) {
      setOpen(true)
      timeoutRef.current = window.setTimeout(() => setOpen(false), 5000)
      return
    }

    if (timeoutRef.current !== null) {
      window.clearTimeout(timeoutRef.current)
    }

    setOpen(false)

    if (props.onClick) {
      props.onClick()
    }
  }

  return <IconButton
    className={classNames(
      props.className,
      styles.DeleteForeverButton,
      open && styles.DeleteForeverButton_danger,
    )}
    onClick={handleClick}
  >
    {open ? <DeleteForever /> : <Delete />}
  </IconButton>
})
