import styled from "styled-components";

function Navbar() {
  return (
    <NavbarSection>
      <Container>
        <div aria-hidden="true">
          <img src="/assets/dappnode_logo.png" alt="logo" height={50} />
          <Inter700>Pinner DAO</Inter700>
        </div>
      </Container>
    </NavbarSection>
  );
}

const NavbarSection = styled.section`
  font-family: Roboto, Helvetica, sans-serif;
  background-color: transparent;
  height: 72px;
  @media only screen and (max-width: 768px) {
    height: auto;
  }
  text-align: center;
  overflow: hidden;
  position: relative;
  z-index: 1;
`;

const Container = styled.div`
  width: 95%;
  margin: auto;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  flex-direction: row;
  @media only screen and (max-width: 768px) {
    flex-direction: column;
  }
  h1 {
    font-family: Roboto, Helvetica, sans-serif;
    font-size: 24px;
    display: flex;
    align-items: center;
    letter-spacing: 0.2px;
    margin-left: 20px;
    img {
      padding-right: 32px;
    }
  }
  div {
    display: flex;
    align-items: center;
    justify-content: space-between;
    flex-direction: row;
    cursor: pointer;
    p {
      margin: 0 15px;
    }
  }
`;

const Inter700 = styled.h1`
  font-family: "Inter-Bold";
  font-weight: bold;
  font-size: 16px;
  line-height: 24px;
  align-items: center;
  color: #35403f, 100%;
  margin-right: 10px;
  img {
    padding-right: 5px;
  }
  &.large {
    font-size: 24px;
  }
`;

export default Navbar;
