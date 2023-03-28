import classNames from "classnames"
import { DropzoneOptions, useDropzone } from "react-dropzone"
import s from "./Dropzone.module.css"

type Props = { className?: string } & DropzoneOptions

export function Dropzone(props: Props) {
  const { className, ...dropzoneOptions} = props

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
    <p>Drop files here</p>
  </div>
}