import React from 'react';
import {MDBContainer, MDBCol, MDBRow, MDBCard, MDBCardText} from 'mdbreact';
import {ProfileMidTabs} from '../components/advanced/ProfileMidTabs';

export const User = () => {
  return (
    <React.Fragment>
      <MDBContainer>
        <MDBRow>
          <MDBCol sm="12" className="py-3">
            <MDBRow>
              <MDBCol sm="12">
                <MDBCard className="card-body p-4 round">
                  <div className="user-image-name d-flex flex-row">
                    <img
                      src="https://mdbootstrap.com/img/Photos/Avatars/avatar-1.jpg"
                      className="rounded-circle float-left mr-4 z-depth-1"
                      alt="aligment"
                    />
                    <div className="d-flex flex-column my-3">
                      <h4>Username</h4>
                      <p className="small">エンジニア</p>
                    </div>
                  </div>
                  <MDBCardText className="mt-4 mb-4">
                    Some quick example text to build on the panel title and make
                    up the bulk of the panel's content.
                  </MDBCardText>
                </MDBCard>
                <ProfileMidTabs />
              </MDBCol>
            </MDBRow>
          </MDBCol>
        </MDBRow>
      </MDBContainer>
    </React.Fragment>
  );
};
