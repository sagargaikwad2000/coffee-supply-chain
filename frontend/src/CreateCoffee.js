import React from 'react';
import { Form, Input, Button, DatePicker, notification, Card } from 'antd';
import axios from 'axios';

const { TextArea } = Input;

const CreateCoffee = () => {
  const [form] = Form.useForm();

  const onFinish = async (values) => {
    try {
      await axios.post('http://localhost:3000/coffeebean/create', values, {
        headers: {
          'Accept': '*/*',
          'Content-Type': 'application/json',
        },
      });
      notification.success({
        message: 'Success',
        description: 'Coffee bean created successfully!',
      });
      form.resetFields();
    } catch (error) {
      notification.error({
        message: 'Error',
        description: 'Failed to create coffee bean.',
      });
    }
  };

  return (
    <div style={{ padding: '20px' }}>
      <Card title="Create Coffee Bean" style={{ maxWidth: '600px', margin: '0 auto' }}>
        <Form
          form={form}
          layout="vertical"
          onFinish={onFinish}
        >
          <Form.Item
            label="Coffee Bean ID"
            name="coffeebeanId"
            rules={[{ required: true, message: 'Please input the coffee bean ID!' }]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            label="Name"
            name="name"
            rules={[{ required: true, message: 'Please input the coffee bean name!' }]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            label="Type"
            name="type"
            rules={[{ required: true, message: 'Please input the coffee bean type!' }]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            label="Quantity Produced"
            name="qtyProduced"
            rules={[{ required: true, message: 'Please input the quantity produced!' }]}
          >
            <Input type="number" />
          </Form.Item>
          <Form.Item
            label="Date Produced"
            name="dateProduced"
            rules={[{ required: true, message: 'Please select the date produced!' }]}
          >
            <DatePicker format="YYYY-MM-DD" />
          </Form.Item>
          <Form.Item
            label="Produced By"
            name="producedBy"
            rules={[{ required: true, message: 'Please input the producer!' }]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            label="Location"
            name="location"
            rules={[{ required: true, message: 'Please input the location!' }]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            label="Cost Per Kg"
            name="costPerKg"
            rules={[{ required: true, message: 'Please input the cost per kg!' }]}
          >
            <Input type="number" />
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit">
              Submit
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </div>
  );
};

export default CreateCoffee;
