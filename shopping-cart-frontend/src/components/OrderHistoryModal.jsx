import React, { useState } from "react";
import "../styles/orderhistorymodal.css";

const OrderHistoryModal = ({ isOpen, onClose, orders }) => {
  const [expandedOrderIds, setExpandedOrderIds] = useState([]);

  if (!isOpen) return null;

  const toggleOrderItems = (orderId) => {
    setExpandedOrderIds((prev) =>
      prev.includes(orderId)
        ? prev.filter((id) => id !== orderId)
        : [...prev, orderId]
    );
  };

  return (
    <div className="order-modal-overlay">
      <div className="order-modal">
        <button className="close-btn" onClick={onClose}>
          &times;
        </button>
        <h2>Order History</h2>

        {orders.length === 0 ? (
          <p>No previous orders found.</p>
        ) : (
          orders.map((order) => {
            const isExpanded = expandedOrderIds.includes(order.ID);
            return (
              <div key={order.ID} className="order-card">
                <p><strong>Order ID:</strong> {order.ID}</p>
                <p><strong>Date:</strong> {new Date(order.CreatedAt).toLocaleString()}</p>

                <button
                  className="toggle-items-btn"
                  onClick={() => toggleOrderItems(order.ID)}
                >
                  {isExpanded ? "Hide Items" : "Show Items"}
                </button>

                {isExpanded && (
                  <div className="ordered-items-list">
                    {order.cart?.cart_items?.map((cartItem, index) => (
                      <div key={index} className="ordered-item">
                        <img
                          src={cartItem.item?.image}
                          alt={cartItem.item?.name}
                          className="ordered-item-img"
                        />
                        <div className="ordered-item-details">
                          <p><strong>{cartItem.item?.name}</strong></p>
                          <p>Price: ₹{cartItem.item?.price}</p>
                          <p>Qty: {cartItem.quantity}</p>
                        </div>
                      </div>
                    ))}
                  </div>
                )}
              </div>
            );
          })
        )}
      </div>
    </div>
  );
};

export default OrderHistoryModal;
