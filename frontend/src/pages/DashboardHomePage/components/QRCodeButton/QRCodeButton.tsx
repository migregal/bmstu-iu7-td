import React, { useLayoutEffect, useRef, useState } from "react"
import Modal from "react-modal"
import classNames from "classnames"

import { ReactComponent as CloseIcon } from "assets/icons/Close.svg"

import styles from "./QRCodeButton.module.css"
import { ReactComponent as QRCodeIcon } from "./QRCode.svg"
import IconButton from "components/IconButton"
import QRCodeStyling from "qr-code-styling"

type Props = {
    className?: string
    url?: string
}

function QRCodeImage(props: Props) {
  const { url } = props
  const ref = useRef<HTMLDivElement | null>(null)
  const qrCodeRef = useRef<QRCodeStyling | null>(null)
  const [copied, setCopied] = useState(false)

  useLayoutEffect(() => {
    qrCodeRef.current = new QRCodeStyling({
      width: 300,
      height: 300,
      image: "/favicon.ico",
      data: url,
      dotsOptions: {
        color: "#7209B7",
        type: "rounded",
      },
      imageOptions: {
        margin: 10
      },
    })

    if (ref.current?.childElementCount === 0) {
      qrCodeRef.current.append(ref.current!)
      setTimeout(() => ref.current!.style.opacity = "1", 10)

      qrCodeRef.current.getRawData().then(
        blob => {
          if (!blob) {
            console.error("QRCodeImage.getRawData blob is null")
            return
          }
          navigator.clipboard.write([
            new ClipboardItem({
              [blob.type]: blob,
            })
          ])
          setCopied(true)
        }
      ).catch(error => {
        console.error("QRCodeImage.getRawData", error)
      })
    }

  }, [url])

  return <>
    <div className={styles.image} ref={ref} />
    <p className={classNames(styles.copied, copied && styles.copied_done)}>
      Copied to clipboard
    </p>
  </>
}

export const QRCodeButton = React.memo(function QRCodeButton(props: Props) {
  const { className } = props
  const [open, setOpen] = useState(false)

  return <>
    <IconButton
      className={classNames(styles.button, className)}
      onClick={() => setOpen(true)}
      title="QR Code"
    >
      <QRCodeIcon />
    </IconButton>
    {open && (
      <Modal
        isOpen={true}
        onRequestClose={() => setOpen(false)}
        preventScroll
        className={styles.modal}
        overlayClassName={styles.overlay}
      >
        <IconButton
          className={styles.modal__close}
          onClick={() => setOpen(false)}
        >
          <CloseIcon />
        </IconButton>

        <QRCodeImage url={props.url} />

      </Modal>
    )}
  </>
})
