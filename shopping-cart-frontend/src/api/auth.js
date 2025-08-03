import axios from "axios";

const API = axios.create({
  baseURL: "http://localhost:8080", // Adjust if your backend runs on a different port
});

// ✅ Corrected: Accepts a single object with username and password
export const loginUser = async ({ username, password }) => {
  const res = await API.post("/users/login", { username, password });
  return res.data; // { token: "..." }
};
