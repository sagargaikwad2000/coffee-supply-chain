import React, { useState } from 'react';
import './producerHome.css';  // Import the CSS file
import axios from 'axios';

const CreateBatchForm = () => {

  const loggedInUser = localStorage.getItem("user") || "";

  const [newBatch, setNewBatch] = useState({
    docType: "",
    batchId: "",
    coffeeType: "",
    location: "",
    createdOn: "",
    status: "",
    quantity: "",
    costPerKg: "",
    producerId: "",
    inspectorId: "",
    processorId: "",
    exporterId: "",
    importerId: ""
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setNewBatch({
      ...newBatch,
      [name]: value
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    newBatch.producerId = loggedInUser

    axios.post("http://localhost:8080/api/create-batch", newBatch, {
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(function (response) {
      alert("Batch created successfully")
      console.log(response);
    }).catch(function (error) {
      alert("Batch creation failed", error);
      console.log(error);
    });

    setNewBatch({
      docType: "",
      batchId: "",
      coffeeType: "",
      location: "",
      createdOn: "",
      status: "",
      quantity: "",
      costPerKg: "",
      producerId: "",
      inspectorId: "",
      processorId: "",
      exporterId: "",
      importerId: ""
    }); // Reset the form
  };

  return (
    <div className="container">
      <h3>Create New Batch</h3>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Coffee Type:</label>
          <input
            type="text"
            name="coffeeType"
            value={newBatch.coffeeType}
            onChange={handleChange}
            required
          />
        </div>
        <div>
          <label>Location:</label>
          <input
            type="text"
            name="location"
            value={newBatch.location}
            onChange={handleChange}
            required
          />
        </div>
        <div>
          <label>Created On:</label>
          <input
            type="date"
            name="createdOn"
            value={newBatch.createdOn}
            onChange={handleChange}
          />
        </div>
        <div>
          <label>Quantity:</label>
          <input
            type="text"
            name="quantity"
            value={newBatch.quantity}
            onChange={handleChange}
            required
          />
        </div>
        <div>
          <label>Cost Per Kg:</label>
          <input
            type="text"
            name="costPerKg"
            value={newBatch.costPerKg}
            onChange={handleChange}
          />
        </div>
        <button type="submit">Add Batch</button>
      </form>
    </div>
  );
};

export default CreateBatchForm;
