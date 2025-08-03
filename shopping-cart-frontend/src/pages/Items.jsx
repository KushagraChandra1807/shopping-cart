import React, { useEffect, useState } from "react";
import "../styles/items.css";
import { fetchItems } from "../api/items";
import { addToCart, fetchCart } from "../api/cart";
import { checkoutCart, getOrders } from "../api/orders";
import { toast } from "react-toastify";
import CartModal from "../components/CartModal";
import OrderHistoryModal from "../components/OrderHistoryModal";

const Items = () => {
  const [items, setItems] = useState([]);
  const [cart, setCart] = useState([]);
  const [orders, setOrders] = useState([]);
  const [showCart, setShowCart] = useState(false);
  const [showOrders, setShowOrders] = useState(false);

  const token = localStorage.getItem("token");

  useEffect(() => {
    const fetchItemsData = async () => {
      try {
        const fetchedItems = await fetchItems(token);
        setItems(fetchedItems);
      } catch (error) {
        console.error("Error fetching items:", error);
      }
    };

    fetchItemsData();
  }, [token]);

  const handleAddToCart = async (itemId) => {
    try {
      await addToCart(token, itemId);
      toast.success("✅ Product added to cart!");
    } catch (error) {
      toast.error("❌ Failed to add product to cart.");
    }
  };

  const handleViewCart = async () => {
    try {
      const data = await fetchCart(token);
      setCart(data.CartItems || []);
      setShowCart(true);
      setShowOrders(false);
    } catch (error) {
      toast.error("❌ Failed to load cart.");
    }
  };

  const handleCheckout = async () => {
    try {
      await checkoutCart(token);
      toast.success("✅ Checkout Successful");
      setShowCart(false);
    } catch (error) {
      toast.error("❌ Checkout failed.");
    }
  };

  const handleOrderHistory = async () => {
    try {
      const data = await getOrders(token);
      setOrders(data);
      setShowOrders(true);
      setShowCart(false);
    } catch (error) {
      toast.error("❌ Failed to load order history.");
    }
  };

  const getImageUrl = (itemName) => {
    const cleaned = itemName
      .toLowerCase()
      .replace(/\s+/g, "")
      .replace(/[^a-z0-9]/g, "");
    return new URL(`/src/assets/${cleaned}.jpg`, import.meta.url).href;
  };

  return (
    <div className="container">
      <div className="top-buttons">
        <button onClick={handleViewCart}>View Cart</button>
        <button onClick={handleCheckout}>Checkout</button>
        <button onClick={handleOrderHistory}>Order History</button>
      </div>

      <div className="grid">
        {items.map((item) => (
          <div className="card" key={item.ID}>
            <img
              src={getImageUrl(item.name)}
              alt={item.name}
              onError={(e) =>
                (e.target.src = new URL(
                  "/src/assets/placeholder.jpg",
                  import.meta.url
                ).href)
              }
            />
            <h2>{item.name}</h2>
            <p>₹{item.price}</p>
            <button onClick={() => handleAddToCart(item.ID)}>
              Add to Cart
            </button>
          </div>
        ))}
      </div>

      <CartModal
        isOpen={showCart}
        onClose={() => setShowCart(false)}
        cartItems={cart}
        setCartItems={setCart}
        onCheckout={handleCheckout}
      />

      <OrderHistoryModal
        isOpen={showOrders}
        onClose={() => setShowOrders(false)}
        orders={orders}
      />
    </div>
  );
};

export default Items;
