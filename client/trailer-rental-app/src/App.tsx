import "./App.css";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Header from "./components/Header";
import Home from "./pages/Home";
import Trailers from "./pages/Trailers";
import Payment from "./pages/Payment";

function App() {
    return (
        <Router>
            <div>
                <Header />
                <div className="pt-16">
                    <Routes>
                        <Route path="/" element={<Home />} />
                        <Route path="/trailers" element={<Trailers />} />
                        <Route path="/payment/:id" element={<Payment />} />
                    </Routes>
                </div>
            </div>
        </Router>
    );
}

export default App;
