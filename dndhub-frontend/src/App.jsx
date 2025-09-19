import { BrowserRouter, Routes, Route, Link } from "react-router-dom";
import { useState } from "react";
import Login from "./pages/Login";
import Register from "./pages/Register";
import FranchiseList from "./pages/FranchiseList";
import FranchiseDetail from "./pages/FranchiseDetail";

export default function App() {
  const [token, setToken] = useState(localStorage.getItem("token"));

  function handleLogin() {
    setToken(localStorage.getItem("token"));
  }

  function handleLogout() {
    localStorage.removeItem("token");
    setToken(null);
  }

  return (
    <BrowserRouter>
      <nav className="p-2 bg-gray-200 flex gap-4">
        <Link to="/">Franchises</Link>
        {!token && <Link to="/login">Login</Link>}
        {!token && <Link to="/register">Register</Link>}
        {token && <button onClick={handleLogout}>Logout</button>}
      </nav>

      <Routes>
        <Route path="/" element={token ? <FranchiseList /> : <p>Please log in</p>} />
        <Route path="/login" element={<Login onLogin={handleLogin} />} />
        <Route path="/register" element={<Register />} />
        <Route path="/franchise/:id" element={<FranchiseDetail />} />
      </Routes>
    </BrowserRouter>
  );
}
