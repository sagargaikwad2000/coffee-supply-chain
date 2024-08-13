import React from 'react';
import { Card, Col, Row } from 'antd';
import { useNavigate } from 'react-router-dom';

const { Meta } = Card;

const ProducerHome = () => {
  const navigate = useNavigate();

  const handleNavigate = (path) => {
    navigate(path);
  };

  return (
    <div style={{ padding: '20px' }}>
      <Row gutter={24}>
        <Col span={8}>
          <Card
            hoverable
            onClick={() => handleNavigate('/create-coffee')}
            cover={<img alt="Create Coffee" src="https://via.placeholder.com/300" />}
          >
            <Meta title="Create Coffee" description="Navigate to Create Coffee page" />
          </Card>
        </Col>
        <Col span={8}>
          <Card
            hoverable
            onClick={() => handleNavigate('/get-all-coffee')}
            cover={<img alt="Get All Coffee" src="https://via.placeholder.com/300" />}
          >
            <Meta title="Get All Coffee" description="Navigate to Get All Coffee page" />
          </Card>
        </Col>
      </Row>
    </div>
  );
};

export default ProducerHome;
