import axios from "axios";

export const checkoutCart = async (token) => {
  await axios.post(
    "http://localhost:8080/orders",
    {},
    { headers: { Authorization: `Bearer ${token}` } }
  );
};

export const getOrders = async (token) => {
  const res = await axios.get("http://localhost:8080/orders", {
    headers: { Authorization: `Bearer ${token}` },
  });
  return res.data;
};
