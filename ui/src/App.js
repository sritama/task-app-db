import './App.css';
import { BrowserRouter, Routes, Route } from "react-router-dom"
import { About } from "./components/About";
import { Tasks } from "./components/Tasks";
function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route
          path="/"
          element={
            <div className="container">
              <Tasks />
            </div>
          }
        />
        <Route
          path="/about"
          element={
            <div className="container">
              <About />
            </div>
          }
        />
      </Routes>
      {/*<div className="App">*/}
      {/*  <h1>*/}
      {/*    List of tasks to do*/}
      {/*  </h1>*/}

      {/*</div>*/}
    </BrowserRouter>
  );
}

export default App;
