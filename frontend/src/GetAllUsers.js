import React, { useState, useEffect } from 'react';
import axios from 'axios';
import "./getAllUsers.css"

const GetAllUsers = () => {
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    axios.get("http://localhost:8080/api/users").then(function (response) {
      setUsers(response.data)
      console.log(response.data)
    }).catch(function (error) {
      console.log(error);
    });;
  }, []);



  return (
    <div>
      <h2>All Users</h2>
      <div className="table-container">
        <table>
          <thead>
            <tr>
              <th>User ID</th>
              <th>First Name</th>
              <th>Last Name</th>
              <th>Email</th>
              <th>Contact No</th>
              <th>Address</th>
              <th>Role</th>
              <th>Doc Type</th>
            </tr>
          </thead>
          <tbody>
            {users.map((user, index) => (
              <tr key={index}>
                <td>{user.userId}</td>
                <td>{user.firstName}</td>
                <td>{user.lastName}</td>
                <td>{user.email}</td>
                <td>{user.contactNo}</td>
                <td>{user.address}</td>
                <td>{user.role}</td>
                <td>{user.docType}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default GetAllUsers;
