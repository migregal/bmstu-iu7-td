import React from "react"
import Modal from "react-modal"
import ReactDOM from "react-dom/client"
import App from "./App"
import "./index.css"

Modal.setAppElement("#root")

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
)
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
)
