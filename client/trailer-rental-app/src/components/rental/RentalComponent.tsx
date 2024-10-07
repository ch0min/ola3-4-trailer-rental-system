import React, { useRef, useState } from "react";
import { FaRegCalendarAlt, FaMapPin } from "react-icons/fa";
import { useNavigate } from "react-router-dom";

function RentalComponent() {
    const navigate = useNavigate();

    const startTimeRef = useRef<HTMLInputElement | null>(null);
    const endTimeRef = useRef<HTMLInputElement | null>(null);
    const [zipCode, setZipCode] = useState<string>("")

    const handleCalendarClick = (
        inputRef: React.RefObject<HTMLInputElement>
    ) => {
        if (inputRef.current) {
            inputRef.current.showPicker();
        }
    };

    const handleSearch = async (e: React.FormEvent) => {
        e.preventDefault();

        if (zipCode) {
                navigate("/trailers", { state: { zipCode: Number(zipCode) } });
    
        }
    };

    return (
        <div className="max-w-4xl flex flex-col items-center p-8 bg-slate-900 bg-opacity-75 rounded-lg shadow-md">
            <form
                onSubmit={handleSearch}
                className="flex flex-row justify-center items-center space-x-2 w-full"
            >
                <div className="flex flex-col">
                    <label
                        htmlFor="start-time"
                        className="block text-sm font-medium text-left text-white"
                    >
                        ZIP CODE
                    </label>
                    <div className="flex items-center border border-gray-300 rounded-md">
                        <div className="flex items-center justify-center p-2 border-r border-gray-300 cursor-pointer">
                            <FaMapPin className="text-white" size={25} />
                        </div>
                        <input
                            type="number"
                            id="zip-code"
                            placeholder="Type zip code"
                            value={zipCode}
                            onChange={(e) => setZipCode(e.target.value)}
                            className="mt-1 block w-48 h-12 p-2 bg-transparent text-white rounded-md outline-none"
                        />
                    </div>
                </div>

                <div className="flex flex-col">
                    <label
                        htmlFor="start-time"
                        className="block text-sm font-medium text-left text-white"
                    >
                        FROM
                    </label>
                    <div className="flex items-center border border-gray-300 rounded-md">
                        <div
                            className="flex items-center justify-center p-2 border-r border-gray-300 cursor-pointer"
                            onClick={() => handleCalendarClick(startTimeRef)}
                        >
                            <FaRegCalendarAlt
                                className="text-white"
                                size={25}
                            />
                        </div>
                        <input
                            type="date"
                            id="start-time"
                            ref={startTimeRef}
                            className="mt-1 block w-48 h-12 p-2 bg-transparent text-white rounded-md outline-none"
                        />
                    </div>
                </div>

                <div className="flex flex-col">
                    <label
                        htmlFor="end-time"
                        className="block text-sm font-medium text-left text-white"
                    >
                        TO
                    </label>
                    <div className="flex items-center border border-gray-300 rounded-md">
                        <div
                            className="flex items-center justify-center p-2 border-r border-gray-300 cursor-pointer"
                            onClick={() => handleCalendarClick(endTimeRef)}
                        >
                            <FaRegCalendarAlt
                                className="text-white"
                                size={25}
                            />
                        </div>
                        <input
                            type="date"
                            id="end-time"
                            ref={endTimeRef}
                            className="mt-1 block w-48 h-12 p-2 bg-transparent text-white rounded-md outline-none"
                        />
                    </div>
                </div>

                <div className="flex flex-col">
                    <button
                        type="submit"
                        className="mt-5 w-32 h-14 bg-amber-600 text-white rounded-md shadow-md hover:bg-amber-700 focus:outline-none"
                        style={{ border: "none" }}
                    >
                        <p className="text-white text-lg font-bold">SEARCH</p>
                    </button>
                </div>
            </form>
        </div>
    );
}

export default RentalComponent;
