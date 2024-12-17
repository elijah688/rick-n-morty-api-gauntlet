import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import CharacterCards from './components/CharactersCard/CharactersCard';
import CharacterForm from './components/CharacterForm/CharacterForm';

function App() {
  return (

      <Router >
        <Routes>
          <Route path="/" element={<CharacterCards />} />
          <Route path="/character/:id" element={<CharacterForm />} />
          <Route path="/create" element={<CharacterForm />} />
        </Routes>
      </Router>
  );
}

export default App;
