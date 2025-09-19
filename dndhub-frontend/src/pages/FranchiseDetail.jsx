import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { getFranchise, createWorker, updateFranchise, updateWorker } from "../api";
import WorkerItem from "../components/WorkerItem";

export default function FranchiseDetail() {
    const { id } = useParams();
    const [franchise, setFranchise] = useState(null);

    // edit form for franchise
    const [edit, setEdit] = useState(null);

    // worker form state
    const [worker, setWorker] = useState({
        name: "",
        monthly_cost_cents: 0,
        creativity: 0,
        discipline: 0,
        charisma: 0,
        efficiency: 0,
        exploration: 0,
        notes: ""
    });
    const [showWorkerForm, setShowWorkerForm] = useState(false);

    async function load() {
        try {
            const data = await getFranchise(id);
            setFranchise(data);
            setEdit(data); // initialize edit form with current values
        } catch {
            alert("Failed to load franchise");
        }
    }

    useEffect(() => { load(); }, [id]);

    async function handleWorkerSubmit(e) {
        e.preventDefault();
        try {
            await createWorker(id, worker);
            setWorker({
                name: "",
                monthly_cost_cents: 0,
                creativity: 0,
                discipline: 0,
                charisma: 0,
                efficiency: 0,
                exploration: 0,
                notes: ""
            });
            setShowWorkerForm(false);
            load(); // reload franchise with new worker
        } catch {
            alert("Failed to add worker");
        }
    }

    async function handleFranchiseUpdate(e) {
        e.preventDefault();
        try {
            await updateFranchise(id, edit);
            load();
        } catch {
            alert("Failed to update franchise");
        }
    }

    if (!franchise) return <p className="p-4">Loading...</p>;

    return (
        <div className="p-4">
            <h1 className="text-xl font-bold mb-2">{franchise.name}</h1>

            {/* Franchise edit form */}
            <h2 className="font-bold mt-4">Edit Franchise</h2>
            <form onSubmit={handleFranchiseUpdate} className="grid grid-cols-2 gap-4 mt-2">
                <div className="col-span-2">
                    <label className="block text-sm font-medium">Name</label>
                    <input className="border p-2 w-full"
                        value={edit.name || ""}
                        onChange={e => setEdit({ ...edit, name: e.target.value })} />
                </div>

                <div>
                    <label className="block text-sm font-medium">Funds</label>
                    <input type="number" className="border p-2 w-full"
                        value={edit.funds_cents}
                        onChange={e => setEdit({ ...edit, funds_cents: parseInt(e.target.value) })} />
                </div>

                <div>
                    <label className="block text-sm font-medium">Property Value (cents)</label>
                    <input type="number" className="border p-2 w-full"
                        value={edit.property_value_cents}
                        onChange={e => setEdit({ ...edit, property_value_cents: parseInt(e.target.value) })} />
                </div>

                <div>
                    <label className="block text-sm font-medium">Unskilled Workers</label>
                    <input type="number" className="border p-2 w-full"
                        value={edit.unskilled_workers}
                        onChange={e => setEdit({ ...edit, unskilled_workers: parseInt(e.target.value) })} />
                </div>

                <div>
                    <label className="block text-sm font-medium">Lowskilled Workers</label>
                    <input type="number" className="border p-2 w-full"
                        value={edit.lowskilled_workers}
                        onChange={e => setEdit({ ...edit, lowskilled_workers: parseInt(e.target.value) })} />
                </div>

                <div>
                    <label className="block text-sm font-medium">Highskilled Workers</label>
                    <input type="number" className="border p-2 w-full"
                        value={edit.highskilled_workers}
                        onChange={e => setEdit({ ...edit, highskilled_workers: parseInt(e.target.value) })} />
                </div>

                <div>
                    <label className="block text-sm font-medium">Cost Unskilled (cents)</label>
                    <input type="number" className="border p-2 w-full"
                        value={edit.cost_unskilled_cents}
                        onChange={e => setEdit({ ...edit, cost_unskilled_cents: parseInt(e.target.value) })} />
                </div>

                <div>
                    <label className="block text-sm font-medium">Cost Lowskilled (cents)</label>
                    <input type="number" className="border p-2 w-full"
                        value={edit.cost_lowskilled_cents}
                        onChange={e => setEdit({ ...edit, cost_lowskilled_cents: parseInt(e.target.value) })} />
                </div>

                <div>
                    <label className="block text-sm font-medium">Cost Highskilled (cents)</label>
                    <input type="number" className="border p-2 w-full"
                        value={edit.cost_highskilled_cents}
                        onChange={e => setEdit({ ...edit, cost_highskilled_cents: parseInt(e.target.value) })} />
                </div>

                <div>
                    <label className="block text-sm font-medium">Revenue Modifier (bp)</label>
                    <input type="number" className="border p-2 w-full"
                        value={edit.revenue_modifier_bp}
                        onChange={e => setEdit({ ...edit, revenue_modifier_bp: parseInt(e.target.value) })} />
                </div>

                <div>
                    <label className="block text-sm font-medium">Upkeep Modifier (bp)</label>
                    <input type="number" className="border p-2 w-full"
                        value={edit.upkeep_modifier_bp}
                        onChange={e => setEdit({ ...edit, upkeep_modifier_bp: parseInt(e.target.value) })} />
                </div>

                <button className="bg-blue-500 text-white px-4 py-2 col-span-2">
                    Save Franchise
                </button>
            </form>


            {/* Workers list */}
            <h2 className="mt-6 font-bold">Workers</h2>
            <ul className="mb-4">
                {franchise.workers.map(w => (
                    <WorkerItem key={w.id} franchiseId={id} worker={w} onUpdated={load} />
                ))}
            </ul>

            {/* Collapsible worker form */}
            <button
                onClick={() => setShowWorkerForm(!showWorkerForm)}
                className="bg-green-600 text-white px-3 py-1 mb-2"
            >
                {showWorkerForm ? "Cancel" : "Add Worker"}
            </button>

            {showWorkerForm && (
                <form onSubmit={handleWorkerSubmit} className="grid grid-cols-2 gap-4 mt-2">
                    <div className="col-span-2">
                        <label className="block text-sm font-medium">Name</label>
                        <input className="border p-2 w-full"
                            value={worker.name}
                            onChange={e => setWorker({ ...worker, name: e.target.value })} />
                    </div>

                    <div>
                        <label className="block text-sm font-medium">Monthly Cost (cents)</label>
                        <input type="number" className="border p-2 w-full"
                            value={worker.monthly_cost_cents}
                            onChange={e => setWorker({ ...worker, monthly_cost_cents: parseInt(e.target.value) })} />
                    </div>

                    <div>
                        <label className="block text-sm font-medium">Creativity</label>
                        <input type="number" className="border p-2 w-full"
                            value={worker.creativity}
                            onChange={e => setWorker({ ...worker, creativity: parseInt(e.target.value) })} />
                    </div>

                    <div>
                        <label className="block text-sm font-medium">Discipline</label>
                        <input type="number" className="border p-2 w-full"
                            value={worker.discipline}
                            onChange={e => setWorker({ ...worker, discipline: parseInt(e.target.value) })} />
                    </div>

                    <div>
                        <label className="block text-sm font-medium">Charisma</label>
                        <input type="number" className="border p-2 w-full"
                            value={worker.charisma}
                            onChange={e => setWorker({ ...worker, charisma: parseInt(e.target.value) })} />
                    </div>

                    <div>
                        <label className="block text-sm font-medium">Efficiency</label>
                        <input type="number" className="border p-2 w-full"
                            value={worker.efficiency}
                            onChange={e => setWorker({ ...worker, efficiency: parseInt(e.target.value) })} />
                    </div>

                    <div>
                        <label className="block text-sm font-medium">Exploration</label>
                        <input type="number" className="border p-2 w-full"
                            value={worker.exploration}
                            onChange={e => setWorker({ ...worker, exploration: parseInt(e.target.value) })} />
                    </div>

                    <div className="col-span-2">
                        <label className="block text-sm font-medium">Notes</label>
                        <input className="border p-2 w-full"
                            value={worker.notes}
                            onChange={e => setWorker({ ...worker, notes: e.target.value })} />
                    </div>

                    <button className="bg-green-500 text-white px-4 py-2 col-span-2">
                        Save Worker
                    </button>
                </form>
            )}

        </div>
    );
}
