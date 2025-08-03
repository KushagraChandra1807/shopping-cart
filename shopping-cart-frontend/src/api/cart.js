import axios from "axios";

const API = axios.create({ baseURL: "http://localhost:8080" });


export const addToCart = async (token, itemId) => {
  return await API.post(
    "/carts",
    { item_id: itemId, quantity: 1 },
    { headers: { Authorization: `Bearer ${token}` } }
  );
};


export const fetchCart = async (token) => {
  const res = await API.get("/carts", {
    headers: { Authorization: `Bearer ${token}` },
  });
  return res.data;
};


export const updateCartQuantity = async (token, itemId, quantity) => {
  return await API.put(
    "/carts",
    { item_id: itemId, quantity },
    { headers: { Authorization: `Bearer ${token}` } }
  );
};
