import { useState } from "react";
import { login } from "../api";

export default function Login({ onLogin }) {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    async function handleSubmit(e) {
        e.preventDefault();
        try {
            await login(username, password);
            onLogin();
        } catch {
            alert("Login failed");
        }
    }

    return (
        <form onSubmit={handleSubmit} className="p-4 max-w-sm mx-auto">
            <h1 className="text-xl font-bold mb-2">Login</h1>
            <input className="border p-2 w-full mb-2"
                value={username} onChange={e => setUsername(e.target.value)} placeholder="Username" />
            <input className="border p-2 w-full mb-2"
                type="password" value={password} onChange={e => setPassword(e.target.value)} placeholder="Password" />
            <button className="bg-blue-500 text-white px-4 py-2">Login</button>
        </form>
    );
}
