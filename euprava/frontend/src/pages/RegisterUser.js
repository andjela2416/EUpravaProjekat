import React, { useState } from 'react';
import axios from 'axios';

const Register = () => {
  const [formData, setFormData] = useState({
    email: '',
    password: '',
    first_name: '',
    last_name: '',
    phone: '',
    address: '',
    user_type: ''
  });

  const [registrationSuccess, setRegistrationSuccess] = useState(false);
  const [registrationError, setRegistrationError] = useState('');

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:8080/users/register', formData);
      console.log(response.data);
      setRegistrationSuccess(true);
      setRegistrationError('');
    } catch (error) {
      console.error(error);
      setRegistrationSuccess(false);
      setRegistrationError('Registracija nije uspela. Pokušajte ponovo.');
    }
  };

  return (
      <div className="container mt-5">
        <h2>Registruj Korisnika</h2>
        {registrationSuccess && (
            <div className="alert alert-success" role="alert">
              Registracija uspešna!
            </div>
        )}
        {registrationError && (
            <div className="alert alert-danger" role="alert">
              {registrationError}
            </div>
        )}
        <form onSubmit={handleSubmit}>
          <div className="mb-3">
            <label htmlFor="first_name" className="form-label">Ime</label>
            <input type="text" className="form-control" id="first_name" name="first_name" value={formData.first_name} onChange={handleChange} required />
          </div>
          <div className="mb-3">
            <label htmlFor="last_name" className="form-label">Prezime</label>
            <input type="text" className="form-control" id="last_name" name="last_name" value={formData.last_name} onChange={handleChange} required />
          </div>
          <div className="mb-3">
            <label htmlFor="email" className="form-label">Email</label>
            <input type="email" className="form-control" id="email" name="email" value={formData.email} onChange={handleChange} required />
          </div>
          <div className="mb-3">
            <label htmlFor="password" className="form-label">Lozinka</label>
            <input type="password" className="form-control" id="password" name="password" value={formData.password} onChange={handleChange} required />
          </div>
          <div className="mb-3">
            <label htmlFor="phone" className="form-label">Broj telefona</label>
            <input type="tel" className="form-control" id="phone" name="phone" value={formData.phone} onChange={handleChange} required />
          </div>
          <div className="mb-3">
            <label htmlFor="address" className="form-label">Adresa</label>
            <input type="text" className="form-control" id="address" name="address" value={formData.address} onChange={handleChange} required />
          </div>
          <div className="mb-3">
            <label htmlFor="user_type" className="form-label">Tip korisnika</label>
            <select className="form-control" id="user_type" name="user_type" value={formData.user_type} onChange={handleChange} required>
              <option value="">Izaberite tip korisnika</option>
              <option value="STUDENT">Student</option>
              <option value="DOKTOR">Doktor</option>
              <option value="PROFESSOR">Profesor</option>
              <option value="KUVAR">Kuvar</option>
              <option value="ADMIN">Admin</option>
              <option value="DEZURAN">Dežuran</option>
            </select>
          </div>
          <button type="submit" className="btn btn-primary">Registruj Korisnika</button>
        </form>
      </div>
  );
};

export default Register;
