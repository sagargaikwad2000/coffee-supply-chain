import React, { useEffect, useState } from 'react';
import { Layout, Card, Input, Button, Switch, Form } from 'antd';
import { useNavigate } from 'react-router-dom';
import './Login.css'; // Import your custom CSS
import axios from 'axios';

const { Content } = Layout;

function Login() {
  const [isAdmin, setIsAdmin] = useState(false);
  const [user, setUser] = useState();
  const navigate = useNavigate();

  const handleLoginSwitch = (checked) => {
    setIsAdmin(checked);
  };


  const handleLogin = (values) => {
    console.log('Login values:', values);

    if (values.username == "admin" && values.password == "admin") {
      navigate("/adminhome")
      return
    }

    axios.post("http://localhost:8080/api/user-login", {
      userId: values.username,
      password: values.password
    }).then(function (response) {
      console.log(response);
      switch (response.data.role) {
        case "producer":
          navigate('/producerhome');
          break;
        case "inspector":
          navigate('/inspectorhome');
          break;
        case "processor":
          navigate('/processorhome');
          break;
        case "exporter":
          navigate('/exporterhome');
          break;
        case "importer":
          navigate('/importerhome');
          break;
        default:
          navigate("/adminhome")
          break;
      }

    }).catch(function (error) {
      alert("User login failed", error);
      console.log(error);
    });


    // Check credentials
    // if (isAdmin && values.username === 'admin' && values.password === 'admin') {
    //   navigate('/adminhome');
    // } else if (!isAdmin && values.username === 'user' && values.password === 'user') {
    //   navigate('/userhome');
    // } else {
    //   notification.error({
    //     message: 'Login Failed',
    //     description: 'Invalid credentials. Please try again.',
    //   });
    // }
  };

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Content className="login-container">
        <div className="login-image">
          <img src="https://www.gep.com/prod/s3fs-public/styles/blog_hero_banner/public/blog-images/Why-Is-There-a-Global-Coffee-Shortage.jpg.webp?itok=wVkwPu9P" alt="Login" style={{ width: '100%', height: '100%' }} />
        </div>
        <div className="login-form-container">
          <Card title={isAdmin ? "Admin Login" : "User Login"} className="login-card">
            {/* <Switch
              checkedChildren="Admin"
              unCheckedChildren="User"
              onChange={handleLoginSwitch}
              style={{ marginBottom: '16px' }}
            /> */}
            <Form
              name="login"
              initialValues={{ remember: true }}
              onFinish={handleLogin}
            >
              <Form.Item
                name="username"
                rules={[{ required: true, message: 'Please input your username!' }]}
              >
                <Input placeholder="Username" />
              </Form.Item>
              <Form.Item
                name="password"
                rules={[{ required: true, message: 'Please input your password!' }]}
              >
                <Input.Password placeholder="Password" />
              </Form.Item>
              <Form.Item>
                <Button type="primary" htmlType="submit" style={{ width: '100%' }}>
                  Log in
                </Button>
              </Form.Item>
            </Form>
          </Card>
        </div>
      </Content>
    </Layout>
  );
}

export default Login;
