import React from 'react';
import { Card, Col, Row } from 'antd';
import { useNavigate } from 'react-router-dom';
import "./producerHome.css"

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
            cover={<img alt="Create Batch" src="https://via.placeholder.com/300" />}
          >
            <Meta title="Create Batch" />
          </Card>
        </Col>
        <Col span={8}>
          <Card
            hoverable
            onClick={() => handleNavigate('/get-all-coffee')}
            cover={<img alt="Get All Batch" src="https://via.placeholder.com/300" />}
          >
            <Meta title="Get All Batch" />
          </Card>
        </Col>
      </Row>
    </div>
  );
};

export default ProducerHome;
