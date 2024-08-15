import React, { useState } from 'react';
import axios from 'axios'

const ExporterBatchForm = ({ formData, onUpdate }) => {
    const [localFormData, setLocalFormData] = useState(formData);
    const loggedInUser = localStorage.getItem("user") || "";

    localFormData.exporterId = loggedInUser

    const handleChange = (e) => {
        const { name, value } = e.target;
        setLocalFormData({
            ...localFormData,
            [name]: value
        });
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        axios.post("http://localhost:8080/api/update-batch", localFormData, {
            headers: {
                'Content-Type': 'application/json'
            }
        }).then(function (response) {
            alert("Batch updated successfully")
            console.log(response);
        }).catch(function (error) {
            alert("Batch updation failed", error);
            console.log(error);
        });
    };

    return (
        <form onSubmit={handleSubmit}>
            <div>
                <label>Batch ID:</label>
                <input
                    type="text"
                    name="batchId"
                    value={localFormData.batchId}
                    onChange={handleChange}
                    readOnly
                />
            </div>
            <div>
                <label>Location:</label>
                <input
                    type="text"
                    name="location"
                    value={localFormData.location}
                    onChange={handleChange}
                />
            </div>
            <div>
                <label>Created On:</label>
                <input
                    type="date"
                    name="createdOn"
                    value={localFormData.createdOn}
                    onChange={handleChange}
                />
            </div>
            <div>
                <label>Status:</label>
                <input
                    type="text"
                    name="status"
                    value={localFormData.status}
                    onChange={handleChange}
                />
            </div>
            <div>
                <label>Destination:</label>
                <input
                    type="text"
                    name="status"
                    value={localFormData.location}
                    onChange={handleChange}
                />
            </div>
            <div>
                <label>Cost Per Kg:</label>
                <input
                    type="number"
                    name="costPerKg"
                    value={localFormData.costPerKg}
                    onChange={handleChange}
                />
            </div>
            <button type="submit">Submit</button>
        </form>
    );
};

export default ExporterBatchForm;
