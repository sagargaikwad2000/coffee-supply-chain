import React, { useEffect, useState } from 'react';
import InspectorBatchForm from './ProcessorBatchForm';
import './processorHome.css';
import axios from 'axios';


const ProcessorHome = () => {
    const [batches, setBatches] = useState([]);

    useEffect(() => {
        axios.get("http://localhost:8080/api/batches").then(function (response) {
            setBatches(response.data)
            console.log(response.data)
        }).catch(function (error) {
            console.log(error);
        });
    }, [])


    const [editingBatch, setEditingBatch] = useState(null);

    const handleUpdate = (updatedBatch) => {
        setBatches(batches.map(batch => batch.batchId === updatedBatch.batchId ? updatedBatch : batch));
        setEditingBatch(null);
    };

    const handleEditClick = (batch) => {
        if (batch.status === "Created") {
            setEditingBatch(batch);
        } else {
            alert("This batch cannot be edited because its status is not 'Created'.");
        }
    };

    return (
        <div className="container">
            <h2>Batch List</h2>
            {batches.length > 0 ?
                <div>
                    <table>
                        <thead>
                            <tr>
                                <th>Batch ID</th>
                                <th>Coffee Type</th>
                                <th>Location</th>
                                <th>Created On</th>
                                <th>Status</th>
                                <th>Quantity</th>
                                <th>Cost Per Kg</th>

                                <th>Actions</th>
                            </tr>
                        </thead>
                        <tbody>
                            {batches.map((batch) => (
                                <tr key={batch.batchId}>
                                    <td>{batch.batchId}</td>
                                    <td>{batch.coffeeType}</td>
                                    <td>{batch.location}</td>
                                    <td>{batch.createdOn || 'N/A'}</td>
                                    <td>{batch.status}</td>
                                    <td>{batch.quantity}</td>
                                    <td>{batch.costPerKg || 'N/A'}</td>
                                    <td>
                                        <button
                                            onClick={() => handleEditClick(batch)}
                                            disabled={batch.status !== "Created"}
                                        >
                                            Update
                                        </button>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>

                    {editingBatch && (
                        <div>
                            <h3>Update Batch Details</h3>
                            <InspectorBatchForm formData={editingBatch} onUpdate={handleUpdate} />
                        </div>
                    )}
                </div> : "No Batches Found"

            }
        </div >
    );
};

export default ProcessorHome;
