import React, { useState } from "react"
import classNames from "classnames"

import { ReactComponent as Close} from "assets/icons/Close.svg"
import IconButton from "components/IconButton"

import errorsObjectToArray from "../../helpers/errorsObjectToArray"

import styles from "./EditableTitle.module.css"
import { ReactComponent as Edit} from "./Edit.svg"
import { ReactComponent as Save} from "./Save.svg"

type Props = {
    href: string
    value: string
    onSave?: (newValue: string) => Promise<{ errors?: Record<string, string> | null }>
    className?: string
}

export const EditableTitle = React.memo(function EditableTitle(props: Props) {
  const [editing, setEditing] = useState(false)
  const [value, setValue] = useState(props.value)
  const [errors, setErrors] = useState<null | string[]>(null)
  const [loading, setLoading] = useState(false)

  const handleEdit = () => {
    setValue(props.value)
    setEditing(true)
  }

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { value } = event.target
    setValue(value)
    setErrors(null)
  }

  const handleSave = async () => {
    if (props.value === value) {
      setEditing(false)
      return
    }

    setLoading(true)
    try {
      if (props.onSave) {
        const { errors } = await props.onSave(value)
        if (errors) {
          setErrors(errorsObjectToArray(errors, ["default", "title"]))
        } else {
          setErrors(null)
        }
      }
      setEditing(false)
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } catch (error: any) {
      console.log("Title.handleSave", error)
      setErrors(error.message || "Unknown error")
    } finally {
      setLoading(false)
    }
  }

  if (editing) {
    return <div className={classNames(styles.title, props.className, styles.title_editing)}>
      <div className={styles.title__row}>
        <input
          className={styles.title__input}
          value={value}
          name="title"
          onChange={handleChange}
          disabled={loading}
        />
        <IconButton className={styles.save} onClick={handleSave} disabled={loading}>
          {value === props.value ? <Close /> : <Save />}
        </IconButton>
      </div>
      <ul className={styles.title__errors}>
        {errors?.map((error, i) => <li key={i}>{error}</li>)}
      </ul>
    </div>
  } else {
    return <div className={styles.title}>
      <a href={props.href} target="_blank" rel="noreferrer" >{props.value}</a>
      <IconButton className={styles.title__edit} onClick={handleEdit}>
        <Edit />
      </IconButton>
    </div>
  }
})
