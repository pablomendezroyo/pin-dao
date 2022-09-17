import React, { useState, useEffect } from "react";
import "./App.css";
import Navbar from "./components/Navbar/Navbar";
import DomainsTable from "./components/DomainsTable/DomainsTable";
import getAllEnsHash from "./logic/getAllEnsHash";
import { EnsHash } from "./types/types";
import { CircularProgress } from "@mui/material";

function App() {
  const [ensHashArray, setEnsHashArray] = useState<EnsHash[]>([]);

  useEffect(() => {
    const fetchFunction = async () => {
      const response = await getAllEnsHash();
      setEnsHashArray(response);
    };

    fetchFunction().catch((error) => {
      console.log(error);
    });
  }, []);

  return (
    <div className="App">
      <Navbar />

      {ensHashArray.length > 0 ? (
        <DomainsTable ensHashArray={ensHashArray} />
      ) : (
        <CircularProgress />
      )}
    </div>
  );
}

export default App;
