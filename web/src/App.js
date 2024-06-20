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
import NewCampaignPage from './pages/newCampaignPage';
import NewCampaignUseCase from './domain/usecases/newCampaign.usecase';
import NewCandidatePage from './pages/newCandidatePage';
import NewCandidateUseCase from './domain/usecases/newCandidate.usecase';
import SingleCandidatePage from './pages/singleCandidatePage';
import SingleCandidateUseCase from './domain/usecases/singleCandidate.usecase';

const campaignlistUseCase = new CampaignListUseCase()
const singleCampaignUseCase = new SingleCampaignUseCase();
const newCampaignUseCase = new NewCampaignUseCase();
const newCandidateUseCase = new NewCandidateUseCase();
const singleCandidateUseCase = new SingleCandidateUseCase();

function App() {


  return (

    <BrowserRouter>
      <Routes>
        <Route path='/' element={<Home />} />
        <Route path='/login' element={<Login />} />
        <Route path='/admin' element={<AdminTemplate />}>
          <Route index element={<AdminDashboad />} />
          {/* <Route path="campaigns" element={<PaginationTable idField={"uuid"} endpoint={campaignlistUseCase} exposedFields={['title', 'issueTime', 'expire']} headers={['Campaign Name', 'Issue Time', 'Expires']} title="Campaigns" />} /> */}
          <Route path="campaigns" element={<CampaignListPage usecase={campaignlistUseCase} />} />
          <Route path="campaign/:uuid" element={<SingleCampaignPage usecase={singleCampaignUseCase} />} />
          <Route path="campaign/new" element={<NewCampaignPage usecase={newCampaignUseCase} />} />
          <Route path="campaign/:campaignUUID/new/candidate" element={<NewCandidatePage usecase={newCandidateUseCase} />} />
          <Route path="candidate/:uuid" element={<SingleCandidatePage usecase={singleCandidateUseCase}/>}/>
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
