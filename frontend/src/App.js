import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './Home';
import './App.css';
import AdminHomepage from './AdminHomepage';
import CreateUser from './CreateUser';
import GetAllUsers from './GetAllUsers';
import GetAllCoffee from './GetAllCoffee';
import ProducerHome from './ProducerHome';
import CreateCoffee from './CreateCoffee';
import InspectorcHome from './InspectorcHome';
import ProcessorHome from './ProcessorHome';
import ExporterHome from './ExporterHome';
import ImporterHome from './ImporterHome';

function App() {
  return (
    <Router>
      <div className="App">


        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/adminhome" element={<AdminHomepage />} />
          <Route path="/create-user" element={<CreateUser />} />
          <Route path="/get-all-users" element={<GetAllUsers />} />
          <Route path="/producerhome" element={<ProducerHome />} />
          <Route path="/create-coffee" element={<CreateCoffee />} />
          <Route path="/inspectorhome" element={<InspectorcHome />} />
          <Route path="/processorhome" element={<ProcessorHome />} />
          <Route path="/get-all-coffee" element={<GetAllCoffee />} />
          <Route path="/exporterhome" element={<ExporterHome />} />
          <Route path="/importerhome" element={<ImporterHome />} />





        </Routes>
      </div>
    </Router>
  );
}

export default App;
