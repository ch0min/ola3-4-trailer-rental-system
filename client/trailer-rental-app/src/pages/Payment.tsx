import React from "react";
import { useLocation, useParams } from "react-router-dom";

function Payment() {
    const { id } = useParams<{ id: string }>();
    const location = useLocation();

    const { zipCode, imageUrl } = (location.state as {
        zipCode: number;
        imageUrl: string;
    }) || {
        zipCode: 0,
        imageUrl: "",
    };

    return (
        <div className="flex flex-col items-center p-8">
            <h2 className="text-2xl font-bold mb-4">Payment</h2>
            <form action="" className="flex flex-col space-y-4">
                <div>
                    <label
                        htmlFor="email"
                        className="block text-sm font-medium text-left"
                    >
                        Email
                    </label>
                    <input
                        type="email"
                        id="email"
                        placeholder="Type email"
                        required
                        className="mt-1 block w-72 h-12 p-2 border border-gray-300 rounded-md"
                    />
                </div>
                <div>
                    <img
                        src={imageUrl}
                        alt={`Trailer ${id}`}
                        className="w-full h-40 object-cover rounded-md mb-2"
                    />
                    <p className="text-lg">Trailer ID: {id}</p>
                    <p className="text-lg">Zip Code: {zipCode}</p>
                </div>

                <button
                    type="submit"
                    className="w-78 h-14 bg-amber-600 text-white rounded-md shadow-md hober:bg-amber-700 focus:outline-none"
                >
                    Book Now
                </button>
            </form>
        </div>
    );
}

export default Payment;
