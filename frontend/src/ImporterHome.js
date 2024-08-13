import React, { useEffect, useState } from 'react';
import { Table, Button, notification, Modal, Form, Input, InputNumber } from 'antd';
import axios from 'axios';

const ImporterHome = () => {
  const [data, setData] = useState([]);
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [selectedBeanId, setSelectedBeanId] = useState(null);
  const [form] = Form.useForm();

  useEffect(() => {
    // Fetch exported coffee data from the API
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
          description: 'Failed to fetch exported coffee data.',
        });
      });
  }, []);

  const handleImport = (values) => {
    const { buyerName, quantity } = values;
    axios.post('http://localhost:3000/coffeebean/buy', {
      coffeebeanId: selectedBeanId,
      buyerName,
      quantity,
    })
      .then(() => {
        notification.success({
          message: 'Success',
          description: `Coffee bean with ID ${selectedBeanId} has been imported by ${buyerName}.`,
        });
        setData(prevData => prevData.map(item =>
          item.key === selectedBeanId ? { ...item, status: 'Imported' } : item
        ));
        setIsModalVisible(false);
        form.resetFields();
      })
      .catch(error => {
        console.error('Error importing coffee bean:', error);
        notification.error({
          message: 'Error',
          description: `Failed to import coffee bean with ID ${selectedBeanId}.`,
        });
      });
  };

  const showImportModal = (key) => {
    setSelectedBeanId(key);
    setIsModalVisible(true);
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
      title: 'Destination',
      dataIndex: 'destination',
      key: 'destination',
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
          onClick={() => showImportModal(record.key)}
        >
          Import
        </Button>
      ),
    },
  ];

  return (
    <div>
      <h1>Exported Coffee</h1>
      <Table
        dataSource={data}
        columns={columns}
        rowKey="key"
      />

      <Modal
        title="Import Coffee"
        visible={isModalVisible}
        onCancel={() => setIsModalVisible(false)}
        onOk={() => form.submit()}
      >
        <Form
          form={form}
          layout="vertical"
          onFinish={handleImport}
        >
          <Form.Item
            label="Buyer Name"
            name="buyerName"
            rules={[{ required: true, message: 'Please enter the buyer\'s name' }]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            label="Quantity"
            name="quantity"
            rules={[{ required: true, message: 'Please enter the quantity' }]}
          >
            <InputNumber min={1} />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
};

export default ImporterHome;
