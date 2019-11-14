import React from "react";
import { MDBContainer, MDBCol, MDBRow } from "mdbreact";
import { Navbar } from "../components/advanced/Navbar";
import { Footer } from "../components/advanced/Footer";

export const Login = () => (
  <React.Fragment>
    <Navbar />
    <MDBContainer>
      <MDBRow>
        <MDBCol sm="12">
          <h1>ようこそページ</h1>
        </MDBCol>
      </MDBRow>
    </MDBContainer>
    <Footer />
  </React.Fragment>
);