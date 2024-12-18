import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';

const CharacterForm = () => {
    const { id } = useParams();
    const navigate = useNavigate();

    const [character, setCharacter] = useState({
        id: null,
        name: '',
        status: '',
        species: '',
        type: '',
        gender: '',
        origin: null,
        location: null,
        image: '',
        episodes: [],
    });

    const isEditMode = Boolean(id);

    useEffect(() => {
        if (isEditMode) {
            fetch(`http://localhost:8081/character/${id}`)
                .then((res) => {
                    if (res.status === 404) {
                        return navigate("/not-found")
                    }
                    return res.json()
                })
                .then((data) => setCharacter({
                    ...data,
                    episodes: data.episodes ?? [],

                }))
                .catch((error) => console.error('Failed to fetch character', error));
        }
    }, [id]);

    const handleChange = (e) => {
        const { name, value } = e.target;
        setCharacter((prev) => ({ ...prev, [name]: value }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        try {
            const response = await fetch('http://localhost:8081/character', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(character),
            });

            if (response.ok) {
                if (response == null) {
                    navigate('/not-found');
                    return
                }
                navigate('/');
            } else {
                console.error('Failed to save character');
            }
        } catch (error) {
            console.error('Error:', error);
        }
    };

    return (
        <div className="w-full h-full min-h-screen flex flex-col items-center justify-center bg-gray-50 p-4">


            <h1 className="text-3xl font-bold text-center mb-6">
                {isEditMode ? 'Edit Character' : 'Create Character'}
            </h1>
            <div className="w-full max-w-4xl bg-white rounded-lg shadow-lg overflow-hidden flex flex-col md:flex-row">
                {/* Image Preview */}
                <div className="flex-1 h-64 md:h-auto bg-gray-100 flex justify-center items-center">
                    {character.image ? (
                        <img
                            src={character.image}
                            alt="Character Preview"
                            className="w-full h-full object-contain rounded-md shadow-md"
                        />
                    ) : (
                        <div className="w-full h-full flex items-center justify-center bg-gray-200 text-gray-500 rounded-md">
                            No Image Preview
                        </div>
                    )}
                </div>

                {/* Form */}
                <form
                    onSubmit={handleSubmit}
                    className="flex-1 p-6 space-y-4 flex flex-col justify-between"
                >
                    <div>
                        {/* Name */}
                        <div>
                            <label className="block text-sm font-medium text-gray-700">Name</label>
                            <input
                                type="text"
                                name="name"
                                value={character.name}
                                onChange={handleChange}
                                className="mt-1 block w-full p-2 border border-gray-300 rounded-md"
                                required
                            />
                        </div>
                        {/* Status */}
                        <div>
                            <label className="block text-sm font-medium text-gray-700">Status</label>
                            <input
                                type="text"
                                name="status"
                                value={character.status}
                                onChange={handleChange}
                                className="mt-1 block w-full p-2 border border-gray-300 rounded-md"
                            />
                        </div>
                        {/* Species */}
                        <div>
                            <label className="block text-sm font-medium text-gray-700">Species</label>
                            <input
                                type="text"
                                name="species"
                                value={character.species}
                                onChange={handleChange}
                                className="mt-1 block w-full p-2 border border-gray-300 rounded-md"
                            />
                        </div>
                        {/* Gender */}
                        <div>
                            <label className="block text-sm font-medium text-gray-700">Gender</label>
                            <input
                                type="text"
                                name="gender"
                                value={character.gender}
                                onChange={handleChange}
                                className="mt-1 block w-full p-2 border border-gray-300 rounded-md"
                            />
                        </div>
                    </div>
                    {/* Image */}
                    <div>
                        <label className="block text-sm font-medium text-gray-700">Image URL</label>
                        <input
                            type="url"
                            name="image"
                            value={character.image}
                            onChange={handleChange}
                            className="mt-1 block w-full p-2 border border-gray-300 rounded-md"
                        />
                    </div>
                    {/* Episodes */}
                    <div>
                        <label className="block text-sm font-medium text-gray-700">Episodes</label>
                        <input
                            disabled={true}
                            type="text"
                            name="episodes"
                            value={character.episodes.join(', ')}
                            onChange={(e) =>
                                setCharacter((prev) => ({
                                    ...prev,
                                    episodes: e.target.value.split(',').map((ep) => parseInt(ep.trim(), 10)),
                                }))
                            }
                            className="mt-1 block w-full p-2 border border-gray-300 rounded-md"
                        />
                    </div>
                    {/* Submit Button */}
                    <button
                        type="submit"
                        className="w-full px-4 py-2 bg-blue-500 text-white font-semibold rounded-lg shadow-md hover:bg-blue-600"
                    >
                        {isEditMode ? 'Update Character' : 'Create Character'}
                    </button>
                </form>
            </div>
        </div>
    );
};

export default CharacterForm;
