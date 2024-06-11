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

const campaignlistUse = new CampaignListUseCase()

function App() {


  
  return (
    
      <BrowserRouter>
        <Routes>
          <Route path='/' element={<Home />} />
          <Route path='/login' element={<Login />} />
          <Route path='/admin' element={<AdminTemplate />}>
            <Route index element={<AdminDashboad />} />
            <Route path="campaigns" element={<PaginationTable idField={"uuid"} endpoint={campaignlistUse} exposedFields={['title', 'issueTime', 'expire']} headers={['Campaign Name', 'Issue Time', 'Expires']} title="Campaigns" />} />
            <Route path="campaigns/:uuid" element={<SingleCampaignPage />} />
          </Route>
          {/* <Route path='/admin/campaigns' element={<AdminTemplate renderContent={AdminCampaignsTable}/>}/>
          <Route path='/admin/camoaigns/pending' element={<AdminTemplate renderContent={() => PaginationTable({title: "Pending Campaigns", headers: ['ID', 'Name', 'Salary', 'Contry', 'City']})}/>}/>
          <Route path='/admin/campaigns/new' element={<AdminTemplate />}/>
          <Route path='/admin/campaigns/:uuid' element={<AdminTemplate />}/>

          <Route path='/admin/candidates/campaign/:campaignUUID' element={<AdminTemplate />}/>
          <Route path='/admin/candidates/:uuid' element={<AdminTemplate />}/> */}
          {/* <Route path='/sign/:campaignUUID/:candidateUUID' element={<CandidateSigning />} /> */}
        </Routes>
      </BrowserRouter>
  );
}

export default App;
