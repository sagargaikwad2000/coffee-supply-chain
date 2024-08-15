import React, { useState } from 'react';

const ProcessorBatchForm = ({ formData, onUpdate }) => {
    const [localFormData, setLocalFormData] = useState(formData);

    const handleChange = (e) => {
        const { name, value } = e.target;
        setLocalFormData({
            ...localFormData,
            [name]: value
        });
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        onUpdate(localFormData);
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
                <label>Coffee Type:</label>
                <input
                    type="text"
                    name="coffeeType"
                    value={localFormData.coffeeType}
                    onChange={handleChange}
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
                <label>Quantity:</label>
                <input
                    type="text"
                    name="quantity"
                    value={localFormData.quantity}
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

export default ProcessorBatchForm;
