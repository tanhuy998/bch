import logo from './logo.svg';
import './App.css';
import { BrowserRouter, Routes, Route, Router, Outlet } from 'react-router-dom'
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
import { motion } from "framer-motion";

import './assets/css/animation.css';
import CandidateSigningPage from './pages/candidateSigningPage';

const campaignlistUseCase = new CampaignListUseCase()
const singleCampaignUseCase = new SingleCampaignUseCase();
const newCampaignUseCase = new NewCampaignUseCase();
const newCandidateUseCase = new NewCandidateUseCase();
const singleCandidateUseCase = new SingleCandidateUseCase();

const pageAnimationVariants = {
  hidden: { opacity: 0, x: 0, y: 20 },
  enter: { opacity: 1, x: 0, y: 0 },
  exit: { opacity: 0 },
};

function App() {


  return (

    <BrowserRouter>
      <Routes>
        <Route path='/' element={<Home />} />
        <Route path='/signing' element={<AnimatePage><SingingPageTemplate /></AnimatePage>}>
          <Route path='campaign/:campaignUUID/candidate/:candidateUUID' element={<CandidateSigningPage />} />
        </Route>
        <Route path='/login' element={<AnimatePage><Login /></AnimatePage>} />
        <Route path='/admin' element={<AdminTemplate />}>
          <Route index element={<AnimatePage><AdminDashboad /></AnimatePage>} />
          {/* <Route path="campaigns" element={<PaginationTable idField={"uuid"} endpoint={campaignlistUseCase} exposedFields={['title', 'issueTime', 'expire']} headers={['Campaign Name', 'Issue Time', 'Expires']} title="Campaigns" />} /> */}
          <Route path="campaigns" element={<AnimatePage><CampaignListPage usecase={campaignlistUseCase} /></AnimatePage>} />
          <Route path="campaign/:uuid" element={<AnimatePage><SingleCampaignPage usecase={singleCampaignUseCase} /></AnimatePage>} />
          <Route path="campaign/new" element={<AnimatePage><NewCampaignPage usecase={newCampaignUseCase} /></AnimatePage>} />
          {/* <Route path="campaign/:campaignUUID/new/candidate" element={<NewCandidatePage usecase={newCandidateUseCase} />} /> */}
          <Route path="candidate/:uuid" element={<AnimatePage><SingleCandidatePage usecase={singleCandidateUseCase} /></AnimatePage>} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

function SingingPageTemplate() {

  return (
    <div className="content">
      <div className='container'>
        <div className='row'>
          <br />
          <div className="page-title" style={{ textAlign: 'center' }}>
            <h3>Signing Info</h3>
          </div>
          <Outlet />
          {/* <div className="card">
            
          </div> */}
        </div>
      </div>
    </div>
  )
}

function AnimatePage({ children }) {

  return (
    <motion.div
      initial="hidden"
      animate="enter"
      exit="exit"
      variants={pageAnimationVariants}
      transition={{ duration: 0.2, type: "easeInOut" }}
      className='relative'
    >
      {children}
    </motion.div>
  )
}

export default App;
