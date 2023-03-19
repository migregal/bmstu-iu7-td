import logo from "./logo.svg"
import "./App.css"
import { useState } from "react"

function App() {  
  const [answer, setAnswer] = useState<string | null>(null)

  const handleClick = () => {
    fetch("/api/v1/auth/login", { method: "POST" }).then(
      async (resp) => {
        setAnswer(JSON.stringify(await resp.json(), null, 2))
      }).catch((error) => {
      console.error("App.handleClick", error)
    })
  }

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
        <button onClick={handleClick}>Auth</button>
        <code style={{ whiteSpace: "pre-wrap"}}>{answer}</code>
      </header>
    </div>
  )
}

export default App
