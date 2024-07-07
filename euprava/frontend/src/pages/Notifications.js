import React, { useState, useEffect } from 'react';

const NotificationsForm = () => {
  const [date, setDate] = useState('');
  const [facultyName, setFacultyName] = useState('');
  const [fieldOfStudy, setFieldOfStudy] = useState('');
  const [description, setDescription] = useState('');
  const [appointments, setAppointments] = useState([]);
  const [error, setError] = useState('');
  const [successMessage, setSuccessMessage] = useState('');

  useEffect(() => {
    fetch('http://localhost:8080/notifications')
      .then(response => response.json())
      .then(data => setAppointments(data))
      .catch(error => console.error('Greska u dobavljanju notifikacija:', error));
  }, []);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setSuccessMessage('');

    const appointmentData = {
      date: new Date(date).toISOString(),
      faculty_name: facultyName,
      field_of_study: fieldOfStudy,
      description,
    };

    try {
      const response = await fetch('http://localhost:8080/notifications', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(appointmentData),
      });

      if (!response.ok) {
        throw new Error('Notifikacija je neuspesno poslata.');
      }

      const newAppointment = await response.json();
      setAppointments([...appointments, newAppointment]);
      setSuccessMessage('Notifikacija uspesno poslata.');
    } catch (error) {
      setError(error.message);
    }
  };

  const handleViewAppointment = (id) => {
    window.location.href = `/notifications/${id}`;
  };

  return (
    <div className="container mt-5">
      <h2>Appointment Form</h2>
      {error && <div className="alert alert-danger">{error}</div>}
      {successMessage && <div className="alert alert-success">{successMessage}</div>}
      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <label htmlFor="date" className="form-label">
            Date
          </label>
          <input
            type="datetime-local"
            className="form-control"
            id="date"
            value={date}
            onChange={(e) => setDate(e.target.value)}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="facultyName" className="form-label">
            Faculty Name
          </label>
          <input
            type="text"
            className="form-control"
            id="facultyName"
            value={facultyName}
            onChange={(e) => setFacultyName(e.target.value)}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="fieldOfStudy" className="form-label">
            Field of Study
          </label>
          <input
            type="text"
            className="form-control"
            id="fieldOfStudy"
            value={fieldOfStudy}
            onChange={(e) => setFieldOfStudy(e.target.value)}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="description" className="form-label">
            Description
          </label>
          <textarea
            className="form-control"
            id="description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            required
          />
        </div>
        <button type="submit" className="btn btn-primary">
          Submit
        </button>
      </form>
      <h2 className="mt-5">Appointments</h2>
      <ul className="list-group">
        {appointments.map(appointment => (
          <li key={appointment.id} className="list-group-item">
            <div>
              <strong>{appointment.faculty_name}</strong> - {appointment.field_of_study}
              <p>{appointment.description}</p>
              <button
                className="btn btn-secondary"
                onClick={() => handleViewAppointment(appointment.id)}
              >
                View Appointment
              </button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default NotificationsForm;
