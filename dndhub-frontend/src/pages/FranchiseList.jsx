import { useEffect, useState } from "react";
import { listFranchises, createFranchise } from "../api";
import { Link } from "react-router-dom";

export default function FranchiseList() {
    const [franchises, setFranchises] = useState([]);
    const [name, setName] = useState("");

    async function load() {
        try {
            const data = await listFranchises();
            console.log("Franchise API response:", data); // 👈 add this
            setFranchises(Array.isArray(data) ? data : []);
        } catch (err) {
            console.error("Failed to load franchises", err);
            alert("Failed to load franchises");
        }
    }




    useEffect(() => { load(); }, []);

    async function handleCreate(e) {
        e.preventDefault();
        await createFranchise(name);
        setName("");
        load();
    }

    return (
        <div className="p-4">
            <h1 className="text-xl font-bold mb-2">Your Franchises</h1>
            <ul>
                {Array.isArray(franchises) && franchises.map(f => (
                    <li key={f.id} className="mb-1">
                        <Link className="text-blue-600" to={`/franchise/${f.id}`}>{f.name}</Link>
                    </li>
                ))}
            </ul>

            <form onSubmit={handleCreate} className="mt-4">
                <input className="border p-2 mr-2"
                    value={name} onChange={e => setName(e.target.value)} placeholder="New Franchise" />
                <button className="bg-blue-500 text-white px-3 py-2">Create</button>
            </form>
        </div>
    );
}
