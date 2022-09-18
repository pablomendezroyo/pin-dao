//External components
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
import Button from "@mui/material/Button";

//Types
import { EnsHash, EnsRow } from "../../types/types";

//Themes
import { ThemeProvider } from "@mui/material";
import { headerTheme, cellTheme } from "./DomainsTableStyle";

function createRowData(domain: string, ipfsHash: string) {
  return { domain, ipfsHash };
}

export default function DomainsTable({
  ensHashArray,
}: {
  ensHashArray: EnsHash[];
}) {
  let rows: Array<EnsRow> = [];

  ensHashArray.forEach((ensHash) => {
    rows.push(createRowData(ensHash.domain, ensHash.hash));
  });

  return (
    <TableContainer component={Paper}>
      <Table stickyHeader sx={{ minWidth: 650 }} aria-label="simple table">
        <ThemeProvider theme={headerTheme}>
          <TableHead>
            <TableRow>
              <TableCell align="center">ENS Domain</TableCell>
              <TableCell align="center">IPFS Hash</TableCell>
              <TableCell align="center">Access</TableCell>
            </TableRow>
          </TableHead>
        </ThemeProvider>
        <ThemeProvider theme={cellTheme}>
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
                <TableCell align="center" component="th" scope="row">
                  <Button
                    href={"https://ipfs.io/ipfs/" + row.ipfsHash.substring(6)}
                  >
                    <img
                      src="/assets/open-in-new.svg"
                      alt="open-in-ipfs"
                      height={30}
                    />
                  </Button>
                  <Button
                    href={
                      process.env.REACT_APP_DAPPNODE_URL +
                      "/ipfs/" +
                      row.ipfsHash.substring(6)
                    }
                  >
                    <img
                      src="/assets/dappnode_logo.png"
                      alt="open-in-dappnode"
                      height={30}
                    />
                  </Button>
                  <Button href={"https://" + row.domain.substring(6) + ".limo"}>
                    <img
                      src="/assets/limo.png"
                      alt="open-in-limo"
                      height={30}
                    />
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </ThemeProvider>
      </Table>
    </TableContainer>
  );
}
