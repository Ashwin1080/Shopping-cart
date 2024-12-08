import React, { useState } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import LoginScreen from './pages/LoginScreen';
import SignUpScreen from './pages/SignUpScreen';
import ItemsList from './pages/ItemsList';


const App = () => {
  const [token, setToken] = useState('');

  return (
    <Router>
      <div>
        <Routes>
          <Route
            path="/"
            element={<LoginScreen setToken={setToken} />}
          />
          <Route
            path="/signup"
            element={<SignUpScreen setToken={setToken} />}
          />
          <Route
            path="/items"
            element={token ? <ItemsList token={token} /> : <LoginScreen setToken={setToken} />}
          />
        </Routes>
      </div>
    </Router>
  );
};

export default App;
