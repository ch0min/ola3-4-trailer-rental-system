import React from "react";

function RentalComponent() {
    return (
        <div className="flex flex-col items-center p-6 bg-gray-100 rounded-lg shadow-md">
            <h2 className="text-2xl font-bold mb-4">Book Your Trailer</h2>
            <form action="" className="flex flex-col space-y-4 w-full max-w-md">
                <div>
                    <label
                        htmlFor="start-time"
                        className="block text-sm font-medium text-gray-700"
                    >
                        Start Time
                    </label>
                    <input
                        type="datetime-local"
                        id="start-time"
                        className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm focus:ring focus:Ring-opacity-509"
                    />
                </div>

                <div>
                    <label
                        htmlFor="end-time"
                        className="block text-sm font-medium text-gray-700"
                    >
                        End Time
                    </label>
                    <input
                        type="datetime-local"
                        id="start-time"
                        className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm focus:ring focus:Ring-opacity-509"
                    />
                </div>

                <button
                    type="submit"
                    className="w-full py-2 px-4 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50"
                >
                    Search
                </button>
            </form>
        </div>
    );
}

export default RentalComponent;
