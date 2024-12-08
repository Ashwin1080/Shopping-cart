import React, { useState } from 'react';
import { Link,useNavigate } from 'react-router-dom';
import axios from 'axios';
import '../css/Global.css';

const SignUpScreen = ({ setToken }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const handleSignUp = async () => {
    try {
      const response = await axios.post('http://localhost:8080/users', {
        username,
        password,
      });
      window.alert('Sign Up Successful');
      setToken(response.data.token); 
      navigate('/'); 
    } catch (error) {
      window.alert('Error during sign up');
    }
  };

  return (
    <div className="auth-container">
      <h2>Sign Up</h2>
      <input
        type="text"
        placeholder="Username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />
      <input
        type="password"
        placeholder="Password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />
      <button onClick={handleSignUp}>Sign Up</button>
      <p>Already have an account? <Link to="/">Login</Link></p>
    </div>
  );
};

export default SignUpScreen;
