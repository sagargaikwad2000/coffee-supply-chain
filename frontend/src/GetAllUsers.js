import React, { useState, useEffect } from 'react';
import { Table, Spin, Alert } from 'antd';
import axios from 'axios';

const GetAllUsers = () => {
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    axios.get('http://localhost:3000/users/all', {
      headers: {
        'Accept': '*/*',
      },
    })
      .then((response) => {
        // Log the response to check its format
        console.log("API Response:", response.data);

        try {
          // Check if response.data is a string and parse it
          const data = typeof response.data === 'string' ? JSON.parse(response.data) : response.data;
          
          // Verify if the parsed data is an array
          if (Array.isArray(data)) {
            const transformedData = data.map(item => item.Record);
            setUsers(transformedData);
          } else {
            console.error('Unexpected data format:', data);
          }
        } catch (error) {
          console.error('Error parsing response data:', error);
          setError(error);
        }

        setLoading(false);
      })
      .catch((error) => {
        console.log("==============");
        console.error('Error fetching data:', error);
        setError(error);
        setLoading(false);
      });
  }, []);

  const columns = [
    {
      title: 'User ID',
      dataIndex: 'userId',
      key: 'userId',
    },
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
    },
    {
      title: 'Role',
      dataIndex: 'role',
      key: 'role',
    },
    {
      title: 'Country',
      dataIndex: 'country',
      key: 'country',
    },
  ];

  if (loading) {
    return <Spin size="large" tip="Loading users..." />;
  }

  if (error) {
    return <Alert message="Error" description="Failed to load users." type="error" showIcon />;
  }

  return (
    <div>
      <h2>All Users</h2>
      <Table
        columns={columns}
        dataSource={users}
        rowKey={(record) => record.userId}  // Use unique userId as key
        pagination={false}
      />
    </div>
  );
};

export default GetAllUsers;
