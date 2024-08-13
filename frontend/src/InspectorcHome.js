import React, { useEffect, useState } from 'react';
import { Table, Button, notification } from 'antd';
import axios from 'axios';

const InspectorcHome = () => {
  const [data, setData] = useState([]);

  useEffect(() => {
    // Fetch data from the API
    axios.get("http://localhost:8080/api/assets")
      .then(response => {
        const data = typeof response.data === 'string' ? JSON.parse(response.data) : response.data;
        if (Array.isArray(data)) {
          const transformedData = data.map(item => ({ Key: item.Key, ...item.Record }));
          setData(transformedData);
          console.log("===========");
          console.log(transformedData);
        } else {
          console.error('Unexpected data format:', data);
        }
        console.log('API Response:', response.data); // Check the structure of the response
      })
      .catch(error => {
        console.error('Error fetching data:', error);
        notification.error({
          message: 'Error',
          description: 'Failed to fetch produced coffee data.',
        });
      });
  }, []);

  const handleApprove = (key) => {
    console.log('Approving coffee bean with Key:', key); // Log the key being approved
    axios.put(`http://localhost:3000/coffeebean/approve/${key}`)
      .then(() => {
        notification.success({
          message: 'Success',
          description: `Coffee bean with ID ${key} approved.`,
        });
        // Optionally, refetch data or update state to reflect changes
        setData(prevData => prevData.map(item =>
          item.Key === key ? { ...item, status: 'Approved' } : item
        ));
      })
      .catch(error => {
        console.error('Error approving coffee bean:', error);
        notification.error({
          message: 'Error',
          description: `Failed to approve coffee bean with ID ${key}.`,
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
      title: 'Cost Per Kg',
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
      title: 'Qty Available',
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
          onClick={() => handleApprove(record.Key)}
        >
          Approve
        </Button>
      ),
    },
  ];

  return (
    <div>
      <h1>Produced Coffee</h1>
      <Table
        dataSource={data}
        columns={columns}
        rowKey="Key"
      />
    </div>
  );
};

export default InspectorcHome;
