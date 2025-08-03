import axios from "axios";

const API = axios.create({ baseURL: "http://localhost:8080" });

// Add item to cart (defaults to quantity: 1)
export const addToCart = async (token, itemId) => {
  return await API.post(
    "/carts",
    { item_id: itemId, quantity: 1 },
    { headers: { Authorization: `Bearer ${token}` } }
  );
};

// Get current cart data
export const fetchCart = async (token) => {
  const res = await API.get("/carts", {
    headers: { Authorization: `Bearer ${token}` },
  });
  return res.data;
};

// Update quantity or remove item if quantity is 0
export const updateCartQuantity = async (token, itemId, quantity) => {
  return await API.put(
    "/carts",
    { item_id: itemId, quantity },
    { headers: { Authorization: `Bearer ${token}` } }
  );
};
