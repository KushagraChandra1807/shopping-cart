import axios from "axios";

const API = axios.create({
  baseURL: "http://localhost:8080", 
});


export const loginUser = async ({ username, password }) => {
  const res = await API.post("/users/login", { username, password });
  return res.data; 
};
