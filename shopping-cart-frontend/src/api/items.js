import axios from "axios";

export const fetchItems = async (token) => {
  const res = await axios.get("http://localhost:8080/items", {
    headers: { Authorization: `Bearer ${token}` },
  });
  return res.data;
};
