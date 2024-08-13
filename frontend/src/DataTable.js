import React from 'react';

function DataTable({ data }) {
    return (
        <div>
            <h2>Batch Details</h2>
            <table border="1" style={{ width: '100%', borderCollapse: 'collapse' }}>
                <thead>
                    <tr>
                        <th style={{ padding: '8px', textAlign: 'left' }}>BatchId</th>
                        <th style={{ padding: '8px', textAlign: 'left' }}>CreatedOn</th>
                        <th style={{ padding: '8px', textAlign: 'left' }}>Status</th>
                        <th style={{ padding: '8px', textAlign: 'left' }}>CoffeeFamily</th>
                        <th style={{ padding: '8px', textAlign: 'left' }}>TypeofSeed</th>
                        <th style={{ padding: '8px', textAlign: 'left' }}>FertilizerType</th>
                        <th style={{ padding: '8px', textAlign: 'left' }}>CoffeeVariety</th>
                        <th style={{ padding: '8px', textAlign: 'left' }}>ProcessorName</th>
                        <th style={{ padding: '8px', textAlign: 'left' }}>ShipName</th>

                    </tr>
                </thead>
                <tbody>
                    {Object.keys(data).map((key, index) => (
                        <tr key={index}>
                            <td style={{ padding: '8px', textAlign: 'left' }}>{data[key]}</td>
                            <td style={{ padding: '8px', textAlign: 'left' }}>{data[key]}</td>
                            <td style={{ padding: '8px', textAlign: 'left' }}>{data[key]}</td>
                            <td style={{ padding: '8px', textAlign: 'left' }}>{data[key]}</td>
                            <td style={{ padding: '8px', textAlign: 'left' }}>{data[key]}</td>
                            <td style={{ padding: '8px', textAlign: 'left' }}>{data[key]}</td>
                            <td style={{ padding: '8px', textAlign: 'left' }}>{data[key]}</td>
                            <td style={{ padding: '8px', textAlign: 'left' }}>{data[key]}</td>
                            <td style={{ padding: '8px', textAlign: 'left' }}>{data[key]}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
}

export default DataTable;
