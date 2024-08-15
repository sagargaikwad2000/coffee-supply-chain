import React, { useState } from 'react';
import './createUser.css';
import axios from 'axios';

const CreateUser = () => {
  const [newUser, setNewUser] = useState({
    docType: "",
    userId: "",
    firstName: "",
    lastName: "",
    email: "",
    contactNo: "",
    address: "",
    password: "",
    role: ""
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setNewUser({
      ...newUser,
      [name]: value
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    axios.post("http://localhost:8080/api/create-user", newUser, {
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(function (response) {
      alert("User created successfully")
      console.log(response);
    }).catch(function (error) {
      alert("User creation failed", error);
      console.log(error);
    });

    setNewUser({
      docType: "",
      userId: "",
      firstName: "",
      lastName: "",
      email: "",
      contactNo: "",
      address: "",
      password: "",
      role: ""
    });
  };

  return (
    <div className="user-container">
      <h3>Add New User</h3>
      <form onSubmit={handleSubmit}>
        <div>
          <label>User ID:</label>
          <input
            type="text"
            name="userId"
            value={newUser.userId}
            onChange={handleChange}
            required
          />
        </div>
        <div>
          <label>First Name:</label>
          <input
            type="text"
            name="firstName"
            value={newUser.firstName}
            onChange={handleChange}
            required
          />
        </div>
        <div>
          <label>Last Name:</label>
          <input
            type="text"
            name="lastName"
            value={newUser.lastName}
            onChange={handleChange}
            required
          />
        </div>
        <div>
          <label>Email:</label>
          <input
            type="email"
            name="email"
            value={newUser.email}
            onChange={handleChange}
            required
          />
        </div>
        <div>
          <label>Contact No:</label>
          <input
            type="text"
            name="contactNo"
            value={newUser.contactNo}
            onChange={handleChange}
            required
          />
        </div>
        <div>
          <label>Address:</label>
          <input
            type="text"
            name="address"
            value={newUser.address}
            onChange={handleChange}
            required
          />
        </div>
        <div>
          <label>Password:</label>
          <input
            type="password"
            name="password"
            value={newUser.password}
            onChange={handleChange}
            required
          />
        </div>
        <div>
          <label>Role:</label>
          <select
            name="role"
            value={newUser.role}
            onChange={handleChange}
            required
          >
            <option value="">Select Role</option>
            <option value="producer">Producer</option>
            <option value="inspector">Inspector</option>
            <option value="processor">Processor</option>
            <option value="exporter">Exporter</option>
            <option value="importer">Importer</option>
          </select>
        </div>
        <button type="submit">Add User</button>
      </form>
    </div>
  );
};

export default CreateUser;
