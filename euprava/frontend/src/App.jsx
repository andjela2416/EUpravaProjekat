import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Navbar from './components/Navbar';
import Home from './pages/Home';
import LoginForm from './components/Login';
import RegisterUser from './pages/RegisterUser';
import University from './pages/University';
import Dorms from './pages/Dorms';
import HealthCare from './pages/HealthCare';
import Food from './pages/Food';

function App() {
  return (
    <div>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<LoginForm />} />
        <Route path="/register/user" element={<RegisterUser />} />
        <Route path="/university" element={<University />} />
        <Route path="/dorms" element={<Dorms />} />
        <Route path="/health-care" element={<HealthCare />} />
        <Route path="/food" element={<Food />} />
      </Routes>
    </div>
  );
}

export default App;

