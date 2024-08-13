import React, { useEffect, useState } from 'react';
import { Table, Button, Modal, Input, notification } from 'antd';
import axios from 'axios';

const ExporterHome = () => {
  const [data, setData] = useState([]);
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [currentKey, setCurrentKey] = useState(null);
  const [destination, setDestination] = useState('');

  useEffect(() => {
    // Fetch processed coffee data from the API
    axios.get("http://localhost:8080/api/assets")
      .then(response => {
        const data = typeof response.data === 'string' ? JSON.parse(response.data) : response.data;
        if (Array.isArray(data)) {
          const transformedData = data.map(item => ({
            key: item.Key,
            ...item.Record,
          }));
          setData(transformedData);
        } else {
          console.error('Unexpected data format:', data);
        }
      })
      .catch(error => {
        console.error('Error fetching data:', error);
        notification.error({
          message: 'Error',
          description: 'Failed to fetch processed coffee data.',
        });
      });
  }, []);

  const showModal = (key) => {
    setCurrentKey(key);
    setIsModalVisible(true);
  };

  const handleExport = () => {
    axios.put(`http://localhost:3000/coffeebean/ship/${currentKey}`, { destination })
      .then(() => {
        notification.success({
          message: 'Success',
          description: `Coffee bean with ID ${currentKey} shipped to ${destination}.`,
        });
        setIsModalVisible(false);
        setData(prevData => prevData.map(item =>
          item.key === currentKey ? { ...item, status: 'Shipped' } : item
        ));
      })
      .catch(error => {
        console.error('Error exporting coffee bean:', error);
        notification.error({
          message: 'Error',
          description: `Failed to ship coffee bean with ID ${currentKey}.`,
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
      title: 'Processed Date',
      dataIndex: 'processedDate',
      key: 'processedDate',
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
          onClick={() => showModal(record.key)}
        >
          Export
        </Button>
      ),
    },
  ];

  return (
    <div>
      <h1>Processed Coffee</h1>
      <Table
        dataSource={data}
        columns={columns}
        rowKey="key"
      />
      <Modal
        title="Export Coffee"
        visible={isModalVisible}
        onOk={handleExport}
        onCancel={() => setIsModalVisible(false)}
      >
        <Input
          placeholder="Enter Destination"
          value={destination}
          onChange={(e) => setDestination(e.target.value)}
        />
      </Modal>
    </div>
  );
};

export default ExporterHome;
