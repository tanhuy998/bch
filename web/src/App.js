import logo from './logo.svg';
import './App.css';
import { BrowserRouter, Routes, Route, Router } from 'react-router-dom'
import Home from './pages/home';
import Login from './pages/login'
import AdminHomePage from './pages/adminHomePage';
import CandidateSigning from './pages/candidateSinging';
import AdminTemplate from './components/adminTemplate';
import AdminDashboad from './components/adminDashboard';
import AdminCampaignsTable from './components/adminCampaignTable';
import PaginationTable from './components/paginationTable';
import { Provider } from 'react-redux';

import { createContext } from 'react';
import CampaignList from './api/campaignList.api';
import CampaignListUseCase from './domain/usecases/campaignListUseCase.usecase';
import SingleCampaignPage from './pages/singleCampaignPage';
import SingleCampaignEndPoint from './api/singleCampaign.api';
import SingleCampaignUseCase from './domain/usecases/singleCampaignUseCase.usecase';
import CampaignListPage from './pages/campaignListPage';

const campaignlistUseCase = new CampaignListUseCase()
const singleCampaignUseCase = new SingleCampaignUseCase();

function App() {

  
  return (
    
      <BrowserRouter>
        <Routes>
          <Route path='/' element={<Home />} />
          <Route path='/login' element={<Login />} />
          <Route path='/admin' element={<AdminTemplate />}>
            <Route index element={<AdminDashboad />} />
            {/* <Route path="campaigns" element={<PaginationTable idField={"uuid"} endpoint={campaignlistUseCase} exposedFields={['title', 'issueTime', 'expire']} headers={['Campaign Name', 'Issue Time', 'Expires']} title="Campaigns" />} /> */}
            <Route path="campaigns" element={<CampaignListPage usecase={campaignlistUseCase}/>} />
            <Route path="campaign/:uuid" element={<SingleCampaignPage usecase={singleCampaignUseCase}/>} />
          </Route>
        </Routes>
      </BrowserRouter>
  );
}

export default App;
