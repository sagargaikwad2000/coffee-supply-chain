import React from 'react';
import { Form, Input, Button, Select } from 'antd';
import axios from 'axios';
import { json, useNavigate } from 'react-router-dom';


const { Option } = Select;

const CreateUser = () => {
  const navigate = useNavigate();

  const onFinish = (values) => {
    values["status"] = "Active"
    console.log('Form Values:', values);

    var participant = {
      "user": {
        "userId": values["userId"],
        "firstName": values["firstName"],
        "lastName": values["lastName"],
        "email": values["email"],
        "contactNo": values["contactNo"],
        "address": values["address"],
        "password": values["password"]
      },
      "role": values["role"],
      "status": values["status"]
    }


    // var payload = JSON.stringify(participant)
    // console.log("payload", payload)


    axios.post("http://localhost:8080/api/user-create", participant, {
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(function (response) {
      alert("User created successfully")
      console.log(response);
      navigate("/");
    }).catch(function (error) {
      alert("User creation failed", error);
      console.log(error);
    });


  };

  return (
    <div style={{ maxWidth: '400px', margin: 'auto' }}>
      <h2>Create User</h2>
      <Form
        name="create_user"
        layout="vertical"
        onFinish={onFinish}
        initialValues={{
          userId: '',
          name: '',
          userRole: '',
          country: '',
        }}
      >
        <Form.Item
          label="User ID"
          name="userId"
          rules={[{ required: true, message: 'Please input the User ID!' }]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="First Name"
          name="firstName"
          rules={[{ required: true, message: 'Please input the First name!' }]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Last Name"
          name="lastName"
          rules={[{ required: true, message: 'Please input the Last name!' }]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Email"
          name="email"
          rules={[{ required: true, message: 'Please input the Email!' }]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Contact No."
          name="contactNo"
          rules={[{ required: true, message: 'Please input the Contact no.!' }]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Address"
          name="address"
          rules={[{ required: true, message: 'Please input the Contact no.!' }]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Password"
          name="password"
          rules={[{ required: true, message: 'Please input the Password!' }]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="User Role"
          name="role"
          rules={[{ required: true, message: 'Please select the User Role!' }]}
        >
          <Select placeholder="Select a role">
            <Option value="producer">Producer</Option>
            <Option value="inspector">Inspector</Option>
            <Option value="processor">Processor</Option>
            <Option value="exporter">Exporter</Option>
            <Option value="importer">Importer</Option>
            <Option value="admin">Admin</Option>
          </Select>
        </Form.Item>

        <Form.Item>
          <Button type="primary" htmlType="submit" block>
            Create User
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};

export default CreateUser;
