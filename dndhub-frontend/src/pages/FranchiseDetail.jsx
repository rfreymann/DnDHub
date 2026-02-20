import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { getFranchise, createWorker, updateFranchise, updateWorker, simulateFranchise, getHistory } from "../api";
import WorkerItem from "../components/WorkerItem";

const ACTIVITY_TYPES = [
    { value: "marketing", label: "Marketing (Boost Revenue)" },
    { value: "accounting", label: "Accounting (Reduce Expenses)" },
    { value: "restructuring", label: "Restructuring (Refund Costs)" },
];

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

    // Simulation state
    const [activities, setActivities] = useState([]); // Array of { type: "", worker_id: "" }
    const [simulationResult, setSimulationResult] = useState(null);
    const [simError, setSimError] = useState(null);
    const [history, setHistory] = useState([]);

    async function load() {
        try {
            const data = await getFranchise(id);
            setFranchise(data);
            setEdit(data); // initialize edit form with current values

            const hist = await getHistory(id);
            setHistory(hist || []);
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

    async function handleSimulate() {
        setSimError(null);
        setSimulationResult(null);
        try {
            // Filter out incomplete activities
            const validActivities = activities.filter(a => a.type && a.worker_id);
            const res = await simulateFranchise(id, validActivities);
            setSimulationResult(res);
            load(); // Refresh funds
        } catch (e) {
            setSimError(e.message);
        }
    }

    function addActivity() {
        if (activities.length < 2) {
            setActivities([...activities, { type: "marketing", worker_id: "" }]);
        }
    }

    function removeActivity(index) {
        const newActs = [...activities];
        newActs.splice(index, 1);
        setActivities(newActs);
    }

    function updateActivity(index, field, value) {
        const newActs = [...activities];
        newActs[index] = { ...newActs[index], [field]: value };
        setActivities(newActs);
    }

    // Helper to get available workers for dropdown (filtering is too strict if we hide selected one)
    // Actually, just let them see all, backend validates uniqueness. UI can show simple validation if needed.
    const uniqueWorkers = franchise?.workers || [];

    if (!franchise) return <p className="p-4">Loading...</p>;

    return (
        <div className="p-4 max-w-4xl mx-auto">
            <h1 className="text-3xl font-bold mb-2">{franchise.name}</h1>
            <div className="bg-gray-100 p-4 rounded mb-6 flex justify-between items-center">
                <div>
                    <span className="text-gray-600 block text-sm">Current Funds</span>
                    <span className="text-2xl font-mono font-bold">{franchise.funds_cents} GP</span>
                </div>
                <div>
                    <span className="text-gray-600 block text-sm">Property Value</span>
                    <span className="text-xl font-mono">{franchise.property_value_cents} GP</span>
                </div>
            </div>

            {/* Simulation Section */}
            <div className="border-2 border-indigo-200 bg-indigo-50 p-4 rounded mb-8 shadow-sm">
                <h2 className="text-xl font-bold mb-4 text-indigo-900 border-b border-indigo-200 pb-2">Monthly Simulation</h2>
                <div className="mb-4">
                    <p className="text-sm text-indigo-800 mb-2">Assign workers to perform activities (max 2). Activities impact your monthly profit.</p>

                    {activities.map((act, idx) => (
                        <div key={idx} className="flex gap-2 mb-2 items-center bg-white p-2 rounded shadow-sm">
                            <select
                                className="border p-2 rounded flex-1"
                                value={act.type}
                                onChange={e => updateActivity(idx, 'type', e.target.value)}
                            >
                                {ACTIVITY_TYPES.map(t => <option key={t.value} value={t.value}>{t.label}</option>)}
                            </select>

                            <select
                                className="border p-2 rounded flex-1"
                                value={act.worker_id}
                                onChange={e => updateActivity(idx, 'worker_id', e.target.value)}
                            >
                                <option value="">-- Select Worker --</option>
                                {uniqueWorkers.map(w => (
                                    <option key={w.id} value={w.id}>
                                        {w.name} (Ch:{w.charisma} Ef:{w.efficiency} Di:{w.discipline})
                                    </option>
                                ))}
                            </select>

                            <button onClick={() => removeActivity(idx)} className="text-red-500 font-bold px-2 hover:bg-red-50 rounded">✕</button>
                        </div>
                    ))}

                    {activities.length < 2 && (
                        <button onClick={addActivity} className="text-sm text-indigo-600 font-medium hover:underline">+ Add Activity</button>
                    )}
                </div>

                <div className="flex gap-4 items-center">
                    <button
                        onClick={handleSimulate}
                        className="bg-indigo-600 hover:bg-indigo-700 text-white px-6 py-2 rounded shadow font-bold transition-colors"
                    >
                        Run Monthly Simulation
                    </button>
                    {simError && <span className="text-red-600 font-medium">{simError}</span>}
                </div>

                {simulationResult && (
                    <div className="mt-4 bg-white p-4 rounded border border-green-200 shadow-sm animate-fade-in">
                        <h3 className="font-bold text-green-800 mb-2">Results</h3>
                        <div className="grid grid-cols-2 gap-4 text-sm">
                            <div className="col-span-2 bg-gray-50 p-2 rounded mb-2 border">
                                <span className="block font-bold text-gray-700 border-b mb-1 pb-1">Transparency Log</span>
                                <div className="grid grid-cols-2 gap-x-4 gap-y-1 text-xs">
                                    <span className="text-gray-500">Base Revenue:</span>
                                    <span className="text-right font-mono">{simulationResult.base_revenue}</span>

                                    <span className="text-gray-500">Base Upkeep:</span>
                                    <span className="text-right font-mono">{simulationResult.base_upkeep}</span>

                                    <span className="text-gray-500">Roll (d100):</span>
                                    <span className="text-right font-mono font-bold">{simulationResult.roll}</span>

                                    <span className="text-gray-500">Tier:</span>
                                    <span className="text-right italic">{simulationResult.description}</span>

                                    <span className="text-gray-500">Multiplier:</span>
                                    <span className="text-right font-mono">x{simulationResult.revenue_multiplier}</span>
                                </div>
                            </div>

                            <div>
                                <span className="block text-gray-500">Revenue</span>
                                <span className="font-mono text-green-600">+{simulationResult.revenue} GP</span>
                            </div>
                            <div>
                                <span className="block text-gray-500">Expenses</span>
                                <span className="font-mono text-red-600">-{simulationResult.expenses} GP</span>
                            </div>
                            <div className="col-span-2 border-t pt-2 mt-2">
                                <span className="block text-gray-500 font-bold">Net Profit</span>
                                <span className={`font-mono text-lg font-bold ${simulationResult.profit >= 0 ? 'text-green-700' : 'text-red-700'}`}>
                                    {simulationResult.profit >= 0 ? '+' : ''}{simulationResult.profit} GP
                                </span>
                            </div>
                        </div>
                        {simulationResult.applied_activities?.length > 0 && (
                            <div className="mt-3 text-xs bg-gray-50 p-2 rounded">
                                <p className="font-semibold text-gray-600">Applied Effects:</p>
                                <ul className="list-disc pl-4 text-gray-600">
                                    {simulationResult.applied_activities.map((msg, i) => <li key={i}>{msg}</li>)}
                                </ul>
                            </div>
                        )}
                    </div>
                )}
            </div>



            <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
                <div>
                    {/* Franchise edit form */}
                    <div className="bg-white p-4 rounded shadow-sm border">
                        <h2 className="font-bold border-b pb-2 mb-4 text-lg">Edit Details</h2>
                        <form onSubmit={handleFranchiseUpdate} className="grid grid-cols-2 gap-4">
                            <div className="col-span-2">
                                <label className="block text-xs uppercase font-bold text-gray-500 mb-1">Name</label>
                                <input className="border p-2 w-full rounded focus:ring-2 focus:ring-blue-100 outline-none"
                                    value={edit.name || ""}
                                    onChange={e => setEdit({ ...edit, name: e.target.value })} />
                            </div>

                            <div className="col-span-2"><hr className="my-1" /></div>
                            <div className="col-span-2 text-xs font-bold text-center text-gray-400">METRICS</div>

                            {/* ... Keep the rest of the inputs, just slightly styled for consistency if possible, or leave raw ... */}
                            {/* Trying to keep functionality same but improve layout slightly to match "Where would i implement logic" request implicitly asking for UI */}

                            <div>
                                <label className="block text-sm font-medium">Property Value</label>
                                <input type="number" className="border p-1 w-full rounded"
                                    value={edit.property_value_cents}
                                    onChange={e => setEdit({ ...edit, property_value_cents: parseInt(e.target.value) })} />
                            </div>
                            <div>
                                <label className="block text-xs font-medium text-gray-600">Funds Adjustment</label>
                                <input type="number" className="border p-1 w-full rounded"
                                    value={edit.funds_cents}
                                    onChange={e => setEdit({ ...edit, funds_cents: parseInt(e.target.value) })} />
                            </div>

                            <div className="col-span-2 text-xs font-bold text-center text-gray-400 mt-2">WORKFORCE</div>

                            <div>
                                <label className="block text-xs font-medium text-gray-600">Unskilled Count</label>
                                <input type="number" className="border p-1 w-full rounded"
                                    value={edit.unskilled_workers}
                                    onChange={e => setEdit({ ...edit, unskilled_workers: parseInt(e.target.value) })} />
                            </div>
                            <div>
                                <label className="block text-xs font-medium text-gray-600">Unskilled Cost (GP)</label>
                                <input type="number" className="border p-1 w-full rounded"
                                    value={edit.cost_unskilled_cents}
                                    onChange={e => setEdit({ ...edit, cost_unskilled_cents: parseInt(e.target.value) })} />
                            </div>

                            <div>
                                <label className="block text-xs font-medium text-gray-600">Lowskilled Count</label>
                                <input type="number" className="border p-1 w-full rounded"
                                    value={edit.lowskilled_workers}
                                    onChange={e => setEdit({ ...edit, lowskilled_workers: parseInt(e.target.value) })} />
                            </div>
                            <div>
                                <label className="block text-xs font-medium text-gray-600">Lowskilled Cost (GP)</label>
                                <input type="number" className="border p-1 w-full rounded"
                                    value={edit.cost_lowskilled_cents}
                                    onChange={e => setEdit({ ...edit, cost_lowskilled_cents: parseInt(e.target.value) })} />
                            </div>

                            <div>
                                <label className="block text-xs font-medium text-gray-600">Highskilled Count</label>
                                <input type="number" className="border p-1 w-full rounded"
                                    value={edit.highskilled_workers}
                                    onChange={e => setEdit({ ...edit, highskilled_workers: parseInt(e.target.value) })} />
                            </div>
                            <div>
                                <label className="block text-xs font-medium text-gray-600">Highskilled Cost (GP)</label>
                                <input type="number" className="border p-1 w-full rounded"
                                    value={edit.cost_highskilled_cents}
                                    onChange={e => setEdit({ ...edit, cost_highskilled_cents: parseInt(e.target.value) })} />
                            </div>

                            <div className="col-span-2 text-xs font-bold text-center text-gray-400 mt-2">MODIFIERS</div>

                            <div>
                                <label className="block text-xs font-medium text-gray-600">Revenue Mod (bp)</label>
                                <input type="number" className="border p-1 w-full rounded"
                                    value={edit.revenue_modifier_bp}
                                    onChange={e => setEdit({ ...edit, revenue_modifier_bp: parseInt(e.target.value) })} />
                            </div>

                            <div>
                                <label className="block text-xs font-medium text-gray-600">Upkeep Mod (bp)</label>
                                <input type="number" className="border p-1 w-full rounded"
                                    value={edit.upkeep_modifier_bp}
                                    onChange={e => setEdit({ ...edit, upkeep_modifier_bp: parseInt(e.target.value) })} />
                            </div>

                            <button className="bg-gray-800 hover:bg-black text-white px-4 py-2 col-span-2 rounded mt-2 text-sm font-bold">
                                Save Changes
                            </button>
                        </form>
                    </div>
                </div>

                <div className="bg-white p-4 rounded shadow-sm border h-fit">
                    <div className="flex justify-between items-center mb-4 border-b pb-2">
                        <h2 className="font-bold text-lg">Unique Workers</h2>
                        <button
                            onClick={() => setShowWorkerForm(!showWorkerForm)}
                            className={`px-3 py-1 text-xs font-bold uppercase tracking-wide rounded ${showWorkerForm ? 'bg-red-100 text-red-600' : 'bg-green-100 text-green-600'}`}
                        >
                            {showWorkerForm ? "Cancel" : "+ Add New"}
                        </button>
                    </div>

                    {showWorkerForm && (
                        <form onSubmit={handleWorkerSubmit} className="grid grid-cols-2 gap-2 mb-6 bg-gray-50 p-3 rounded border">
                            <div className="col-span-2">
                                <label className="block text-xs font-bold text-gray-500">Name</label>
                                <input className="border p-1 w-full rounded"
                                    value={worker.name}
                                    onChange={e => setWorker({ ...worker, name: e.target.value })} />
                            </div>

                            <div>
                                <label className="block text-sm font-medium">Monthly Cost (GP)</label>
                                <input type="number" className="border p-1 w-full rounded"
                                    value={worker.monthly_cost_cents}
                                    onChange={e => setWorker({ ...worker, monthly_cost_cents: parseInt(e.target.value) })} />
                            </div>

                            {/* Stats */}
                            {[
                                { l: 'Creativity', k: 'creativity' },
                                { l: 'Discipline', k: 'discipline' },
                                { l: 'Charisma', k: 'charisma' },
                                { l: 'Efficiency', k: 'efficiency' },
                                { l: 'Exploration', k: 'exploration' }
                            ].map(({ l, k }) => (
                                <div key={k}>
                                    <label className="block text-xs font-bold text-gray-500">{l}</label>
                                    <input type="number" className="border p-1 w-full rounded"
                                        value={worker[k]}
                                        onChange={e => setWorker({ ...worker, [k]: parseInt(e.target.value) })} />
                                </div>
                            ))}

                            <div className="col-span-2">
                                <label className="block text-xs font-bold text-gray-500">Notes</label>
                                <input className="border p-1 w-full rounded"
                                    value={worker.notes}
                                    onChange={e => setWorker({ ...worker, notes: e.target.value })} />
                            </div>

                            <button className="bg-green-600 hover:bg-green-700 text-white px-4 py-2 col-span-2 rounded text-sm font-bold">
                                Create Worker
                            </button>
                        </form>
                    )}

                    <ul className="space-y-2">
                        {franchise.workers.map(w => (
                            <WorkerItem key={w.id} franchiseId={id} worker={w} onUpdated={load} />
                        ))}
                        {franchise.workers.length === 0 && <p className="text-gray-400 text-sm text-center italic py-4">No unique workers allowed yet.</p>}
                    </ul>
                </div>

                {/* History Section */}
                <div className="border border-gray-200 bg-white p-4 rounded mb-8 shadow-sm">
                    <h2 className="text-xl font-bold mb-4 text-gray-800 border-b pb-2">History</h2>
                    <div className="max-h-96 overflow-y-auto">
                        <table className="w-full text-sm text-left">
                            <thead className="text-xs text-gray-500 uppercase bg-gray-50 sticky top-0">
                                <tr>
                                    <th className="px-2 py-2">Date</th>
                                    <th className="px-2 py-2 text-right">Rev</th>
                                    <th className="px-2 py-2 text-right">Exp</th>
                                    <th className="px-2 py-2 text-right">Profit</th>
                                    <th className="px-2 py-2 text-center">Roll</th>
                                    <th className="px-2 py-2">Activities</th>
                                </tr>
                            </thead>
                            <tbody>
                                {history.length === 0 ? (
                                    <tr><td colSpan="6" className="text-center py-4 text-gray-400">No history available</td></tr>
                                ) : (
                                    history.map(h => (
                                        <tr key={h.id} className="border-b hover:bg-gray-50">
                                            <td className="px-2 py-2">{new Date(h.created_at).toLocaleDateString()}</td>
                                            <td className="px-2 py-2 text-right font-mono text-green-600">{h.revenue}</td>
                                            <td className="px-2 py-2 text-right font-mono text-red-600">{h.expenses}</td>
                                            <td className={`px-2 py-2 text-right font-mono font-bold ${h.profit >= 0 ? 'text-green-700' : 'text-red-700'}`}>
                                                {h.profit >= 0 ? '+' : ''}{h.profit}
                                            </td>
                                            <td className="px-2 py-2 text-center text-gray-500">{h.roll}</td>
                                            <td className="px-2 py-2 text-xs text-gray-600 italic">{h.activities}</td>
                                        </tr>
                                    ))
                                )}
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    );
}


