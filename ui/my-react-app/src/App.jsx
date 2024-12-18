import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';

import CharacterCards from './components/CharactersCard/CharactersCard';
import CharacterForm from './components/CharacterForm/CharacterForm';
import NotFount from './components/NotFound/NotFound';

function App() {
  return (

      <Router >
        <Routes>
          <Route path="/" element={<CharacterCards />} />
          <Route path="/character/:id" element={<CharacterForm />} />
          <Route path="/create" element={<CharacterForm />} />
          <Route path="/not-found" element={<NotFount />} />
          <Route path="*" element={<Navigate to="/" />} />
        </Routes>
      </Router>
  );
}

export default App;
