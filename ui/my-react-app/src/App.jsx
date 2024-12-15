import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import CharacterCards from './components/CharactersCard/CharactersCard'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
     <CharacterCards></CharacterCards>
    </>
  )
}

export default App
