import React, { useState, useEffect, } from 'react';
import { useNavigate } from "react-router";

const CharacterCards = () => {
  const [characters, setCharacters] = useState([]);
  const [page, setPage] = useState(0);
  const navigate = useNavigate();

  const fetchCharacters = async (page) => {
    const offset = page * 10;
    const response = await fetch(`http://localhost:8081/character?offset=${offset}&limit=4`);
    const data = await response.json();
    setCharacters(data);
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

  useEffect(() => {
    fetchCharacters(page);
  }, [page]);

  const handleNextPage = () => {
    setPage(page + 1);
  };

  const handlePrevPage = () => {
    if (page > 0) {
      setPage(page - 1);
    }
  };

  const handleCreate = () => {
    navigate("/create")
  };


  const handleEdit = (id) => {
    navigate(`/character/${id}`);
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
        <button
          onClick={handleCreate}
          className="px-6 py-2 bg-green-500 text-white font-semibold rounded-lg shadow-md hover:bg-green-600 disabled:bg-green-300"
        >
          Create
        </button>
        <button
          onClick={handleNextPage}
          className="px-6 py-2 bg-blue-500 text-white font-semibold rounded-lg shadow-md hover:bg-blue-600"
        >
          Next
        </button>
      </div>
    </div>
  );
};

export default CharacterCards;
