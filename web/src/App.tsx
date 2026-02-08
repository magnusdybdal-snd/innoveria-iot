import { useState } from 'react'
import reactLogo from './assets/react.svg'
import innLogo from './assets/innoveria.png'
import viteLogo from '/vite.svg'
import './App.css'

function App() {
  const [count, setCount] = useState(0)

  return (
      <>
          <div id="top">
              <ul id="navigation">
                  <li><a href="#" target="_blank">
                      <img src={innLogo} alt="Innoveria logo"/>
                  </a></li>
                  <li>Home</li>
                  <li>Dashboard</li>
                  <li>Sensors</li>
                  <li>Settings</li>
                  <li>
                      <button>
                          Log out
                      </button>
                  </li>
              </ul>
          </div>
          <a padding="2em">
              <div>
                  <a href="https://vite.dev" target="_blank">
                      <img src={viteLogo} className="logo" alt="Vite logo"/>
                  </a>
                  <a href="https://react.dev" target="_blank">
                      <img src={reactLogo} className="logo react" alt="React logo"/>
                  </a>
              </div>
              <h1>Vite + React</h1>
              <div className="card">
                  <button onClick={() => setCount((count) => count + 1)}>
                      count is {count}
                  </button>
                  <p>
                      Edit <code>src/App.tsx</code> and save to test HMR
                  </p>
              </div>
              <p className="read-the-docs">
                  Click on the Vite and React logos to learn more
              </p>
          </a>
      </>
  )
}

export default App
