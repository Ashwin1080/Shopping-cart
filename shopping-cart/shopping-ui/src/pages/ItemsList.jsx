import React, { useEffect, useState } from 'react';
import axios from 'axios';
import Header from './Header';
import '../css/Global.css';

const ItemsList = ({ token }) => {
  const [items, setItems] = useState([]);

  const fetchItems = async () => {
    const response = await axios.get('http://localhost:8080/items', {
      headers: { Authorization: `Bearer ${token}` },
    });
    setItems(response.data);
  };

  const addToCart = async (itemId) => {
    try {
      await axios.post(
        'http://localhost:8080/carts',
        { item_id: itemId },
        { headers: { Authorization: `Bearer ${token}` } }
      );
      window.alert('Item added to cart');
    } catch {
      window.alert('Failed to add item');
    }
  };

  useEffect(() => {
    fetchItems();
  }, []);

  return (
    <div className="items-list">
      <Header token={token} />
      <h2>Items</h2>
      <ul>
        {items.map((item) => (
          <li key={item.id}>
            {item.name} <button onClick={() => addToCart(item.id)}>Add to Cart</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ItemsList;
