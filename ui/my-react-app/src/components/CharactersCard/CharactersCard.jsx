import React, { useState, useEffect } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';

const CharacterCards = () => {
  const [characters, setCharacters] = useState([]);
  const [page, setPage] = useState(0);
  const [inputPage, setInputPage] = useState(0); // Track input field value
  const navigate = useNavigate();
  const limit = 4

  const fetchCharacters = async (page) => {
    const offset = page * limit;
    const response = await fetch(`http://localhost:8081/character?offset=${offset}&limit=${limit}`);
    const data = await response.json();
    if (data === null) {
      navigate('/not-found');
      return
    }
    setCharacters(data);
  };

  useEffect(() => {
    fetchCharacters(page);
  }, [page]);

  const handleNextPage = () => {
    setPage(page + 1);
    setInputPage(page + 1);
  };

  const handlePrevPage = () => {
    if (page > 0) {
      setPage(page - 1);
      setInputPage(page - 1);
    }
  };

  const handleEdit = (id) => {
    navigate(`/character/${id}`);
  };

  const deleteCharacter = async (id) => {
    try {
      const response = await fetch(`http://localhost:8081/character/${id}`, {
        method: 'DELETE',
      });

      if (response.ok) {
        // Update the state to remove the deleted character
        setCharacters((prev) => prev.filter((character) => character.id !== id));
      } else {
        console.error('Failed to delete character');
      }
    } catch (error) {
      console.error('Error:', error);
    }
  };

  const handleInputChange = (e) => {
    const newPage = Number(e.target.value);
    setInputPage(newPage);
  };

  const handleGoToPage = () => {
    setPage(inputPage);
  };

  const handleCreate = () => {
    navigate("/create");
  };

  return (
    <div className="w-full h-full min-h-screen flex flex-col items-center justify-center bg-gray-50 p-4">

      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-2 gap-6">
        {characters.map((character) => (
          <div key={character.id} className="bg-white rounded-lg shadow-lg overflow-hidden">
            <img
              src={character.image}
              alt={character.name}
              className="w-full h-56 object-cover"
            />
            <div className="p-4">
              <h3 className="text-xl font-semibold text-gray-800">{character.name}</h3>
              <p className="text-gray-500">Status: {character.status}</p>
              <p className="text-gray-500">Species: {character.species}</p>
              <p className="text-gray-500">Gender: {character.gender}</p>
              <p className="text-gray-500">Origin: {character?.origin.name}</p>
              <p className="text-gray-500">Current Location: {character?.location.name}</p>
            </div>
            <div className="p-4 flex justify-between">
              <button
                onClick={() => deleteCharacter(character.id)}
                className="px-4 py-2 bg-red-500 text-white font-semibold rounded-lg shadow-md hover:bg-red-600"
              >
                Delete
              </button>
              <button
                onClick={() => handleEdit(character.id)}
                className="px-4 py-2 bg-blue-500 text-white font-semibold rounded-lg shadow-md hover:bg-blue-600"
              >
                Edit
              </button>
            </div>
          </div>
        ))}
      </div>

      {/* Pagination buttons */}
      <div className="flex justify-between items-center mt-8">
        <button
          onClick={handlePrevPage}
          disabled={page === 0}
          className="px-6 py-2 bg-blue-500 text-white font-semibold rounded-lg shadow-md hover:bg-blue-600 disabled:bg-blue-300"
        >
          Prev
        </button>

        {/* Page number input */}
        <div className="flex items-center">
          <input
            type="number"
            value={inputPage}
            onChange={handleInputChange}
            className="px-3 py-2 border border-gray-300 rounded-md"
            min={0}
          />
          <button
            onClick={handleGoToPage}
            className="ml-2 px-4 py-2 bg-blue-500 text-white font-semibold rounded-lg shadow-md hover:bg-blue-600"
          >
            Go to
          </button>
        </div>

        <button
          onClick={handleNextPage}
          className="px-6 py-2 bg-blue-500 text-white font-semibold rounded-lg shadow-md hover:bg-blue-600 disabled:bg-blue-300"
        >
          Next
        </button>
      </div>

      {/* Create Character button */}
      <div className="mt-4">
        <button
          onClick={handleCreate}
          className="px-6 py-2 bg-green-500 text-white font-semibold rounded-lg shadow-md hover:bg-green-600"
        >
          Create Character
        </button>
      </div>
    </div>
  );
};


export default CharacterCards;
