import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
import Button from "@mui/material/Button";
import { EnsHash } from "../../App";

interface ENSRow {
  domain: string;
  ipfsHash: string;
}

function createRowData(domain: string, ipfsHash: string) {
  return { domain, ipfsHash };
}

export default function DomainsTable({
  ensHashArray,
}: {
  ensHashArray: EnsHash[];
}) {
  let rows: Array<ENSRow> = [];

  /*rows = [
    createRowData(
      "hi.pinnerdao.eth",
      "QmYwAPJzv5CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG"
    ),
    createRowData(
      "test.pinnerdao.eth",
      "QmYwAPJzv5CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG"
    ),
  ];*/

  ensHashArray.forEach((ensHash) => {
    rows.push(createRowData(ensHash.domain, ensHash.hash));
  });

  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell align="center"> ENS Domain</TableCell>
            <TableCell align="center"> IPFS Hash</TableCell>
            <TableCell align="center"> Access</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {rows.map((row) => (
            <TableRow
              key={row.domain}
              sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
            >
              <TableCell align="center" component="th" scope="row">
                {row.domain}
              </TableCell>
              <TableCell align="center" component="th" scope="row">
                {row.ipfsHash}
              </TableCell>
              <TableCell align="center">
                <Button>
                  <img
                    src="/assets/open-in-new.svg"
                    alt="open-in-new"
                    height={30}
                  />
                </Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
