import React, { useState, useEffect } from "react";
import "./App.css";
import Navbar from "./Navbar";
import DomainsTable from "./components/DomainsTable/DomainsTable";
import getAllEnsHash from "./logic/getAllEnsHash";
import { EnsHash } from "./types/types";

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
      <DomainsTable ensHashArray={ensHashArray} />
    </div>
  );
}

export default App;
