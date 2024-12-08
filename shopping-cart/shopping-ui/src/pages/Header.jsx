import React from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import '../css/Global.css';

const Header = ({ token }) => {
  const navigate = useNavigate();

  const handleCartClick = async () => {
    const response = await axios.get('http://localhost:8080/carts', {
      headers: { Authorization: `Bearer ${token}` },
    });
    const cartItems = response.data;
    window.alert(JSON.stringify(cartItems));
  };

  const handleOrderHistoryClick = async () => {
    const response = await axios.get('http://localhost:8080/orders', {
      headers: { Authorization: `Bearer ${token}` },
    });
    const orders = response.data;
    window.alert(JSON.stringify(orders));
  };

  const handleCheckout = async () => {
    try {
      await axios.post(
        'http://localhost:8080/orders',
        {},
        { headers: { Authorization: `Bearer ${token}` } }
      );
      window.alert('Order successful');
      navigate('/items'); // Redirect back to items screen after checkout
    } catch {
      window.alert('Failed to place order');
    }
  };

  return (
    <div className="header">
      <button onClick={handleCheckout}>Checkout</button>
      <button onClick={handleCartClick}>Cart</button>
      <button onClick={handleOrderHistoryClick}>Order History</button>
    </div>
  );
};

export default Header;
