import React from "react";
import Header from "../components/Header";
import RentalComponent from "../components/rental/RentalComponent";

function Home() {
    return (
        <div className="flex flex-col min-h-screen bg-gray-50">
            <Header />
            <div className="flex-grow flex justify-center items-center">
                <RentalComponent />
            </div>
        </div>
    );
}

export default Home;
