import React, { useState, useEffect } from 'react';
import { Table, Button, notification } from 'antd';
import axios from 'axios';
import moment from 'moment';

function ProcessorHome() {
    const [coffeebeans, setCoffeebeans] = useState([]);

    useEffect(() => {
        // Fetch approved coffee beans from the API
        axios.get("http://localhost:8080/api/assets")
            .then(response => {
                const data = typeof response.data === 'string' ? JSON.parse(response.data) : response.data;

                if (Array.isArray(data)) {
                    // Include 'Key' in the transformed data
                    const transformedData = data.map(item => ({
                        key: item.Key,  // Include Key here
                        ...item.Record
                    }));
                    setCoffeebeans(transformedData);
                } else {
                    console.error('Unexpected data format:', data);
                }

                console.log(response.data);
            })
            .catch(error => {
                console.log(error);
                notification.error({
                    message: 'Error',
                    description: 'Failed to fetch approved coffee beans.',
                });
            });
    }, []);

    const handleProcess = (coffeebeanId) => {
        const processedDate = moment().format('YYYY-MM-DD'); // Get the current date

        axios.put(`http://localhost:3000/coffeebean/process/${coffeebeanId}`, {
            processedDate: processedDate
        })
            .then(response => {
                notification.success({
                    message: 'Success',
                    description: `Coffee bean ${coffeebeanId} processed successfully.`,
                });

                // Optionally, you can refresh the table data after processing
                setCoffeebeans(prevBeans => prevBeans.filter(bean => bean.key !== coffeebeanId));
            })
            .catch(error => {
                console.error('Error processing coffee bean:', error);
                notification.error({
                    message: 'Error',
                    description: `Failed to process coffee bean ${coffeebeanId}.`,
                });
            });
    };

    const columns = [
        {
            title: 'Name',
            dataIndex: 'name',
            key: 'name',
        },
        {
            title: 'Type',
            dataIndex: 'type',
            key: 'type',
        },
        {
            title: 'Cost per Kg',
            dataIndex: 'costPerKg',
            key: 'costPerKg',
        },
        {
            title: 'Date Produced',
            dataIndex: 'dateProduced',
            key: 'dateProduced',
        },
        {
            title: 'Location',
            dataIndex: 'location',
            key: 'location',
        },
        {
            title: 'Produced By',
            dataIndex: 'producedBy',
            key: 'producedBy',
        },
        {
            title: 'Quantity Available',
            dataIndex: 'qtyAvailable',
            key: 'qtyAvailable',
        },
        {
            title: 'Status',
            dataIndex: 'status',
            key: 'status',
        },
        {
            title: 'Action',
            key: 'action',
            render: (_, record) => (
                <Button
                    type="primary"
                    onClick={() => handleProcess(record.key)}  // Use record.key instead of record.Key
                >
                    Process
                </Button>
            ),
        },
    ];

    return (
        <div style={{ padding: '20px' }}>
            <Table columns={columns} dataSource={coffeebeans} />
        </div>
    );
}

export default ProcessorHome;
