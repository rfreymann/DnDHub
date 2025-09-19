import { useState } from "react";
import { register } from "../api";

export default function Register() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    async function handleSubmit(e) {
        e.preventDefault();
        try {
            await register(username, password);
            alert("Registered! Now log in.");
        } catch {
            alert("Registration failed");
        }
    }

    return (
        <form onSubmit={handleSubmit} className="p-4 max-w-sm mx-auto">
            <h1 className="text-xl font-bold mb-2">Register</h1>
            <input className="border p-2 w-full mb-2"
                value={username} onChange={e => setUsername(e.target.value)} placeholder="Username" />
            <input className="border p-2 w-full mb-2"
                type="password" value={password} onChange={e => setPassword(e.target.value)} placeholder="Password" />
            <button className="bg-green-500 text-white px-4 py-2">Register</button>
        </form>
    );
}
