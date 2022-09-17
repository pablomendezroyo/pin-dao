import "./App.css";
import Navbar from "./Navbar";
import DomainsTable from "./components/DomainsTable/DomainsTable";
import resolveEnsContentHash from "./logic/resolveEnsContentHash";

export interface EnsHash {
  domain: string;
  hash: string;
}

/*const ensHashArray = [
  {domain: "example.pinnerdao.eth", hash: "QmYwAPJzv5CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG"},
  {domain: "test.pinnerdao.eth", hash: "QmYwAPJzv5CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG"}
]*/

function App() {
  return (
    <div className="App">
      <Navbar />
      <DomainsTable ensHashArray={[]} />
    </div>
  );
}

export default App;
