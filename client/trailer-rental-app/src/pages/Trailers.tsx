import React, { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import axios from "axios";
import t1 from "../assets/images/trailer-dummy-1.jpeg";
// import t2 from "../assets/images/trailer-dummy-2.jpeg";
// import t3 from "../assets/images/trailer-dummy-3.jpeg";

interface Trailer {
    id: number;
    name: string;
    price: string;
    zipCode: number;
    imageUrl: string;
}

interface LocationState {
    zipCode: number;
}

function Trailers() {
    const navigate = useNavigate();
    const location = useLocation();

    const { zipCode } = (location.state as LocationState) || { zipCode: 0 };
    const [trailers, setTrailers] = useState<Trailer[]>([])

    useEffect(() => {
        if (zipCode) {
            axios.get(`/api/trailer/${zipCode}`)
            .then(response => setTrailers(response.data))
            .catch(error => console.error("Error fetching trailers:", error))
        }
    }, [zipCode])

    console.log(trailers)

    // const trailers: Trailer[] = [
    //     {
    //         id: 1,
    //         name: "Trailer 1",
    //         price: "$50/day",
    //         zipCode: 2770,
    //         imageUrl: t1,
    //     },
    //     {
    //         id: 2,
    //         name: "Trailer 2",
    //         price: "$60/day",
    //         zipCode: 2770,
    //         imageUrl: t2,
    //     },
    //     {
    //         id: 3,
    //         name: "Trailer 3",
    //         price: "$55/day",
    //         zipCode: 2300,
    //         imageUrl: t3,
    //     },
    // ];

    // const availableTrailers = trailers.filter(
    //     (trailer) => trailer.zipCode === zipCode 
    // );

    const handleContinue = (id: number, imageUrl: string) => {
        navigate(`/payment/${id}`, { state: { zipCode, imageUrl } });
    };

    return (
        <div className="max-w-4xl p-12 bg-transparent bg-opacity-75 rounded-lg">
            <h2 className="text-2xl font-bold text-black mb-4">
                Available Trailers
            </h2>
     
                <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                    {Array.isArray(trailers) && trailers.length > 0 ? (
                trailers.map((trailer) => (
                    <div
                        key={trailer.id}
                        className="p-4 bg-white rounded-md shadow-md"
                    >
                        <img
                            src={trailer.imageUrl || t1}
                            alt={trailer.name}
                            className="w-full h-40 object-cover rounded-md mb-2"
                        />
                        <h3 className="text-lg font-semibold">
                            {trailer.name}
                        </h3>
                        <p className="text-gray-600">{trailer.price}</p>
                        <button
                            onClick={() =>
                                handleContinue(trailer.id, trailer.imageUrl)
                            }
                            className="mt-2 w-full bg-amber-600 text-white rounded-md py-2 hover:bg-amber-700"
                        >
                            Continue
                        </button>
                    </div>
                ))
            ) : (
                <p>No trailers available for this zip code.</p>
            )}
            </div>
        </div>
    );
}

export default Trailers;
