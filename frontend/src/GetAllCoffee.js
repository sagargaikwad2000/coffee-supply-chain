import React, { useState, useEffect } from 'react';
import { Table, Spin, Alert } from 'antd';
import axios from 'axios';

const GetAllCoffee = () => {
  const [coffees, setCoffees] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    axios.get("http://localhost:8080/api/assets").then((response) => {
      console.log("API Response:", response.data);
      const data = typeof response.data === 'string' ? JSON.parse(response.data) : response.data;


      // Check if response.data is an array
      if (Array.isArray(data)) {
        // Map the data to extract the Record part
        const transformedData = data.map(item => item.Record);
        setCoffees(transformedData);
      } else {
        console.error('Unexpected data format:', response.data);
      }
      setLoading(false);
    })
      .catch((error) => {
        console.error('Error fetching data:', error);
        setError(error);
        setLoading(false);
      });
  }, []);

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
      title: 'Produced By',
      dataIndex: 'producedBy',
      key: 'producedBy',
    },
    {
      title: 'Buyer Name',
      dataIndex: 'buyerName',
      key: 'buyerName',
    },
    {
      title: 'Quantity Sold',
      dataIndex: 'quantitySold',
      key: 'quantitySold',
    },
    {
      title: 'Processed Date',
      dataIndex: 'processedDate',
      key: 'processedDate',
    },
    {
      title: 'Destination',
      dataIndex: 'destination',
      key: 'destination',
    },
  ];

  if (loading) {
    return <Spin size="large" tip="Loading coffee beans..." />;
  }

  if (error) {
    return <Alert message="Error" description="Failed to load coffee beans." type="error" showIcon />;
  }

  return (
    <div>
      <h2>All Coffee Beans</h2>
      <Table
        columns={columns}
        dataSource={coffees}
        rowKey={(record) => record}  // Ensure each row has a unique key
        pagination={false}
      />
    </div>
  );
};

export default GetAllCoffee;
