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


function App() {
  return (
    
      <BrowserRouter>
        <Routes>
          <Route path='/' element={<Home />} />
          <Route path='/login' element={<Login />} />
          <Route path='/admin' element={<AdminTemplate />}>
            <Route index element={<AdminDashboad/>}/>
            <Route path="campaigns" element={<PaginationTable headers={['ID', 'Name', 'Salary', 'Contry', 'City']} title="Campaigns" />} />
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
