import React, {useState} from 'react';
import { HashRouter as Router, Routes, Route } from "react-router-dom";
import Index from './views';
import Page from './views/page';
import Search from './views/search';
import {message, Progress} from "antd";

const App: React.FC = () => {
  let [progress, setProgress] = useState(100)
  window.electron.onProgress((event, value)=>{
    setProgress(value)
  })
  return (
    <Router>
      <Progress style={{display:progress==100?'none':'block'}} percent={progress} />
      <Routes>
        <Route path="/" element={<Index />} />
        <Route path="/view/:name/:id/:view" element={<Page />} />
        <Route path="/search" element={<Search />} />
      </Routes>
    </Router>
  );
}

export default App
