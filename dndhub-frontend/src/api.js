const API_URL = "http://localhost:8080"; // your Go backend

function getAuthHeader() {
    const token = localStorage.getItem("token");
    return token ? { Authorization: `Bearer ${token}` } : {};
}

export async function register(username, password) {
    const res = await fetch(`${API_URL}/users/register`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password }),
    });
    if (!res.ok) throw new Error("Registration failed");
    return res.json();
}

export async function login(username, password) {
    const res = await fetch(`${API_URL}/users/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password }),
    });
    if (!res.ok) throw new Error("Login failed");
    const data = await res.json();
    localStorage.setItem("token", data.token);
    return data;
}

export async function listFranchises() {
    const res = await fetch(`${API_URL}/franchise/`, {
        headers: { ...getAuthHeader() },
    });
    if (!res.ok) throw new Error("Failed to load franchises");
    return res.json();
}

export async function createFranchise(name) {
    const res = await fetch(`${API_URL}/franchise/create`, {
        method: "POST",
        headers: { "Content-Type": "application/json", ...getAuthHeader() },
        body: JSON.stringify({ name }),
    });
    if (!res.ok) throw new Error("Failed to create franchise");
    return res.json();
}

export async function getFranchise(id) {
    const res = await fetch(`${API_URL}/franchise/${id}`, {
        headers: { ...getAuthHeader() },
    });
    if (!res.ok) throw new Error("Failed to load franchise");
    return res.json();
}

export async function updateFranchise(id, franchise) {
    const res = await fetch(`${API_URL}/franchise/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json", ...getAuthHeader() },
        body: JSON.stringify(franchise),
    });
    if (!res.ok) throw new Error("Failed to update franchise");
    return res.json();
}

export async function createWorker(franchiseId, worker) {
    const res = await fetch(`${API_URL}/franchise/${franchiseId}/workers`, {
        method: "POST",
        headers: { "Content-Type": "application/json", ...getAuthHeader() },
        body: JSON.stringify(worker),
    });
    if (!res.ok) throw new Error("Failed to create worker");
    return res.json();
}

export async function updateWorker(franchiseId, workerId, worker) {
    const res = await fetch(`${API_URL}/franchise/${franchiseId}/workers/${workerId}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json", ...getAuthHeader() },
        body: JSON.stringify(worker),
    });
    if (!res.ok) throw new Error("Failed to update worker");
    return res.json();
}

export async function deleteWorker(franchiseId, workerId) {
    const res = await fetch(`${API_URL}/franchise/${franchiseId}/workers/${workerId}`, {
        method: "DELETE",
        headers: { ...getAuthHeader() },
    });
    if (!res.ok) throw new Error("Failed to delete worker");
}

