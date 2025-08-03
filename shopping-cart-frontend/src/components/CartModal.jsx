import React from "react";
import "../styles/CartModal.css";
import { updateCartQuantity } from "../api/cart";
import toast from "react-hot-toast";

const CartModal = ({ isOpen, onClose, cartItems, setCartItems, onCheckout }) => {
  const token = localStorage.getItem("token");

  const handleQuantityChange = async (itemId, newQuantity) => {
    if (newQuantity < 0) return;

    try {
      await updateCartQuantity(token, itemId, newQuantity);

      setCartItems((prev) =>
        newQuantity === 0
          ? prev.filter((item) => item.Item.ID !== itemId)
          : prev.map((item) =>
              item.Item.ID === itemId ? { ...item, Quantity: newQuantity } : item
            )
      );
    } catch (err) {
      console.error("Error updating cart quantity:", err);
      toast.error("Failed to update cart.");
    }
  };

  const getSubtotal = () => {
    return cartItems?.reduce((acc, item) => {
      const price = item?.Item?.price || 0;
      const qty = item?.Quantity || 0;
      return acc + price * qty;
    }, 0).toFixed(2);
  };

  const getImageUrl = (itemName) => {
    try {
      const cleaned = itemName
        .toLowerCase()
        .replace(/\s+/g, "")
        .replace(/[^a-z0-9]/g, "");
      return new URL(`/src/assets/${cleaned}.jpg`, import.meta.url).href;
    } catch {
      return new URL("/src/assets/placeholder.jpg", import.meta.url).href;
    }
  };

  if (!isOpen) return null;

  return (
    <div className="cart-modal-overlay">
      <div className="cart-modal">
        <button className="close-btn" onClick={onClose}>
          &times;
        </button>
        <h2>Your Cart</h2>
        {cartItems.length === 0 ? (
          <p className="empty-text">Your cart is empty.</p>
        ) : (
          <div className="cart-items">
            {cartItems.map((cartItem) => {
              const { ID, name, price } = cartItem.Item;
              const quantity = cartItem.Quantity;

              return (
                <div key={ID} className="cart-item">
                  <img
                    src={getImageUrl(name)}
                    alt={name}
                    className="cart-item-img"
                    onError={(e) =>
                      (e.target.src = new URL("/src/assets/placeholder.jpg", import.meta.url).href)
                    }
                  />
                  <div className="cart-item-details">
                    <h4>{name}</h4>
                    <p>Price: ₹{price.toFixed(2)}</p>
                    <div className="qty-controls">
                      <button
                        className="qty-btn"
                        onClick={() => handleQuantityChange(ID, quantity - 1)}
                      >
                        −
                      </button>
                      <span>{quantity}</span>
                      <button
                        className="qty-btn"
                        onClick={() => handleQuantityChange(ID, quantity + 1)}
                      >
                        +
                      </button>
                    </div>
                    <p className="total-line">
                      Total: ₹{(price * quantity).toFixed(2)}
                    </p>
                  </div>
                </div>
              );
            })}

            <div className="subtotal">
              <strong>Subtotal: ₹{getSubtotal()}</strong>
            </div>

            <button className="checkout-btn" onClick={onCheckout}>
              Checkout
            </button>
          </div>
        )}
      </div>
    </div>
  );
};

export default CartModal;
