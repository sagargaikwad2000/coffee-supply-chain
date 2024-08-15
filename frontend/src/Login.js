import React, { useState } from 'react';
import './login.css';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';


const Login = () => {

    const navigate = useNavigate()

    const [credentials, setCredentials] = useState({
        userId: "",
        password: ""
    });

    const handleChange = (e) => {
        const { name, value } = e.target;
        setCredentials({
            ...credentials,
            [name]: value
        });
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        if (credentials.userId == "admin" && credentials.password == "admin") {
            navigate("/adminhome")
            return
        }

        axios.post("http://localhost:8080/api/user-login", {
            userId: credentials.userId,
            password: credentials.password
        }).then(function (response) {
            console.log(response);
            localStorage.setItem("user", credentials.userId)
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

    };

    return (
        <div className="login-container">
            <h3>Login</h3>
            <form onSubmit={handleSubmit}>
                <div>
                    <label>User ID:</label>
                    <input
                        type="text"
                        name="userId"
                        value={credentials.userId}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div>
                    <label>Password:</label>
                    <input
                        type="password"
                        name="password"
                        value={credentials.password}
                        onChange={handleChange}
                        required
                    />
                </div>
                <button type="submit">Login</button>
            </form>
        </div>
    );
};

export default Login;
