import React from 'react';
import { Card, Row, Col } from 'antd';
import { useNavigate } from 'react-router-dom'; // Import useNavigate instead of useHistory

const { Meta } = Card;

function AdminHomepage() {
  const navigate = useNavigate(); // Initialize useNavigate

  const navigateTo = (path) => {
    navigate(path); // Use navigate instead of history.push
  };

  return (
    <div style={{ padding: '20px' }}>
      <Row gutter={16}>
        <Col span={6}>
          <Card
            title="Create User"
            onClick={() => navigateTo('/create-user')}
            hoverable
          >
            <Meta description="Create a new user." />
          </Card>
        </Col>
        <Col span={6}>
          <Card
            title="Get All Users"
            onClick={() => navigateTo('/get-all-users')}
            hoverable
          >
            <Meta description="View a list of all users." />
          </Card>
        </Col>
        <Col span={6}>
          <Card
            title="Get All Coffee"
            onClick={() => navigateTo('/get-all-coffee')}
            hoverable
          >
            <Meta description="View all coffee products." />
          </Card>
        </Col>
        <Col span={6}>
          <Card
            title="Additional Action"
            onClick={() => navigateTo('/additional-action')}
            hoverable
          >
            <Meta description="Another action to be defined." />
          </Card>
        </Col>
      </Row>
    </div>
  );
}

export default AdminHomepage;
