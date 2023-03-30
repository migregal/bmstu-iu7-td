import classNames from "classnames"
import React from "react"
import { DropzoneOptions, useDropzone } from "react-dropzone"
import s from "./Dropzone.module.css"

type Props = { className?: string, title?: React.ReactNode } & DropzoneOptions

export const Dropzone = React.memo(function Dropzone(props: Props) {
  const { className, title, ...dropzoneOptions} = props

  const {
    getInputProps,
    getRootProps,
    isFocused,
    isDragAccept,
    isDragReject,
  } = useDropzone(dropzoneOptions)

  return <div {...getRootProps({className: classNames(s.root, className, {
    [s.root_accept]: isDragAccept,
    [s.root_focus]: isFocused,
    [s.root_reject]: isDragReject,
  })})}>
    <input {...getInputProps()} />
    <p>{title}</p>
  </div>
})
