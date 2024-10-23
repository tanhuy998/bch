import logo from './logo.svg';
import './App.css';
import { BrowserRouter, Routes, Route, Router, Outlet } from 'react-router-dom'
import Home from './pages/home';
import Login from './pages/login/login.page'
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
import CandidateSigningPage from './pages/candidateSinging/candidateSigning.page';
import CandidateSigningUseCase from './domain/usecases/candidateSigning.usecase';
import './config/debug';
import './assets/css/background.css';
import PageNotFound from './pages/404';
import InternalErrorPage from './pages/500';
import AppContext from './contexts/app.context';
import EditSingleCandidatePage from './pages/editSingleCandidatePage';
import EditSingleCandidateUseCase from './domain/usecases/editSingleCandidate.usecase';
import EditSingleCampaignPage from './pages/editSingleCampaignPage';
import EditSingleCampaignUseCase from './domain/usecases/editSingleCampaign.usecase';
import LoginUseCase from './pages/login/usecase'
import SwithcTenantPage from './pages/switchTenant/switchTenant.page';
import SwitchTenantUseCase from './pages/switchTenant/usecase';
// import CandidateSigningPage from './pages/candidateSigning/candidateSinging.page';

const campaignlistUseCase = new CampaignListUseCase()
const singleCampaignUseCase = new SingleCampaignUseCase();
const newCampaignUseCase = new NewCampaignUseCase();
const newCandidateUseCase = new NewCandidateUseCase();
const singleCandidateUseCase = new SingleCandidateUseCase();
const candidateSigningUseCase = new CandidateSigningUseCase();
const editSingleCandidateUseCase = new EditSingleCandidateUseCase();
const editSingleCampaignUsecase = new EditSingleCampaignUseCase();
const loginUseCase = new LoginUseCase();
const switchTenantUseCase = new SwitchTenantUseCase()

const pageAnimationVariants = {
  hidden: { opacity: 0, x: 0, y: 20 },
  enter: { opacity: 1, x: 0, y: 0 },
  exit: { opacity: 0 },
};

function App() {


  return (
    <AppContext.Provider value={{}}>
      <BrowserRouter>
        <Routes>
          <Route path="/404" element={<PageNotFound />} />
          <Route path='/500' element={<InternalErrorPage />} />
          <Route path='/' element={<Home />} />
          <Route path='/signing' element={<AnimatePage><SingingPageTemplate /></AnimatePage>}>
            <Route path='campaign/:campaignUUID/candidate/:candidateUUID' element={<CandidateSigningPage usecase={candidateSigningUseCase} />} />
          </Route>
          <Route path='/login' element={<AnimatePage><Login usecase={loginUseCase}/></AnimatePage>} />
          <Route path="/auth">
            <Route path="switch" element={<AnimatePage><SwithcTenantPage usecase={switchTenantUseCase}/></AnimatePage>} />
          </Route>
          <Route path='/admin' element={<AdminTemplate />}>
            <Route index element={<AnimatePage><AdminDashboad /></AnimatePage>} />
            {/* <Route path="campaigns" element={<PaginationTable idField={"uuid"} endpoint={campaignlistUseCase} exposedFields={['title', 'issueTime', 'expire']} headers={['Campaign Name', 'Issue Time', 'Expires']} title="Campaigns" />} /> */}
            <Route path="campaigns" element={<AnimatePage><CampaignListPage usecase={campaignlistUseCase} /></AnimatePage>} />
            <Route path="campaign/:uuid" element={<AnimatePage><SingleCampaignPage usecase={singleCampaignUseCase} /></AnimatePage>} />
            <Route path="campaign/edit/:campaignUUID" element={<AnimatePage><EditSingleCampaignPage usecase={editSingleCampaignUsecase}/></AnimatePage>} />
            <Route path="campaign/new" element={<AnimatePage><NewCampaignPage usecase={newCampaignUseCase} /></AnimatePage>} />
            {/* <Route path="campaign/:campaignUUID/new/candidate" element={<NewCandidatePage usecase={newCandidateUseCase} />} /> */}

            <Route path="candidate/:uuid" element={<AnimatePage><SingleCandidatePage usecase={singleCandidateUseCase} /></AnimatePage>} />
            <Route path='candidate/edit/:candidateUUID' element={<AnimatePage><EditSingleCandidatePage usecase={editSingleCandidateUseCase} /></AnimatePage>} />
          </Route>
        </Routes>
      </BrowserRouter>
    </AppContext.Provider>
  );
}

function SingingPageTemplate() {

  return (
    <div className="content" style={{
      //backgroundImage: "url(https://vneid.gov.vn/_next/static/media/background-login.98683067.png)",
      backgroundPosition: "50%",
      backgroundSize: "cover",
      paddingTop: "32px",
      backgroundRepeat: "no-repeat",
      width: "100vw",
      height: "100vh",
      filter: "contrast(400%)",
      filter: "saturate(100%)",
      //boxShadow: "0 0 200px rgba(0, 0, 0, 0.5) inset",
    }}>
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
