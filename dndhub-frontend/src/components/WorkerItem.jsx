import { useState } from "react";
import { updateWorker, deleteWorker } from "../api";

export default function WorkerItem({ franchiseId, worker, onUpdated }) {
    const [editing, setEditing] = useState(false);
    const [edit, setEdit] = useState(worker);

    async function handleSave(e) {
        e.preventDefault();
        try {
            await updateWorker(franchiseId, worker.id, edit);
            setEditing(false);
            onUpdated();
        } catch {
            alert("Failed to update worker");
        }
    }

    async function handleDelete() {
        if (!window.confirm(`Delete worker "${worker.name}"?`)) return;
        try {
            await deleteWorker(franchiseId, worker.id);
            onUpdated();
        } catch {
            alert("Failed to delete worker");
        }
    }

    if (!editing) {
        return (
            <li className="mb-2 flex justify-between items-center">
                <span>
                    <span className="font-medium">{worker.name}</span>
                    {" "}
                    (Charisma {worker.charisma}, Efficiency {worker.efficiency})
                </span>
                <span className="flex gap-2">
                    <button
                        className="text-sm text-blue-600"
                        onClick={() => setEditing(true)}
                    >
                        Edit
                    </button>
                    <button
                        className="text-sm text-red-600"
                        onClick={handleDelete}
                    >
                        Delete
                    </button>
                </span>
            </li>
        );
    }

    return (
        <li className="mb-2 border p-2">
            <form onSubmit={handleSave} className="grid grid-cols-2 gap-2">
                {/* inputs as before */}
                <div className="col-span-2 flex gap-2">
                    <button className="bg-blue-500 text-white px-3 py-1">Save</button>
                    <button
                        type="button"
                        className="bg-gray-300 px-3 py-1"
                        onClick={() => setEditing(false)}
                    >
                        Cancel
                    </button>
                </div>
            </form>
        </li>
    );
}
