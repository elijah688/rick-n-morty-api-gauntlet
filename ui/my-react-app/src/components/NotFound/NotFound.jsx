import React from 'react';
import { Link } from 'react-router-dom';

const NotFound = () => {
  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="text-center">
        <h1 className="text-9xl font-extrabold text-blue-600">404</h1>
        <p className="text-2xl text-gray-700 mt-4">Not Found</p>
        <Link
          to="/"
          className="mt-6 inline-block px-6 py-3 text-white bg-blue-600 font-semibold rounded-lg shadow-md hover:bg-blue-700"
        >
          Go Back to Home
        </Link>
      </div>
    </div>
  );
};

export default NotFound;
