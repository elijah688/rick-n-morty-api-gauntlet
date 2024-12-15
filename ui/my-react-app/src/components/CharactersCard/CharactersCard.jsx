import React, { useState, useEffect } from 'react';

const CharacterCards = () => {
  const [characters, setCharacters] = useState([]);
  const [page, setPage] = useState(0);

  // Fetch data from API
  const fetchCharacters = async (page) => {
    const offset = page * 10;
    const response = await fetch(`http://localhost:8081/character?offset=${offset}&limit=10`);
    const data = await response.json();
    setCharacters(data);
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

  return (
    <div className="max-w-screen-xl mx-auto p-4">
      <h1 className="text-3xl font-bold text-center mb-8">Character Cards</h1>

      {/* Render Character Cards in 2 columns of 5 rows */}
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

