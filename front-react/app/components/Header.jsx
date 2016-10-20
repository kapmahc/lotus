import React, { PropTypes } from 'react'
import { connect } from 'react-redux'
import {Navbar, Nav, NavItem} from 'react-bootstrap'
import {Link} from 'react-router'
import {LinkContainer} from 'react-router-bootstrap'

import SwitchLang from './SwitchLang'
import PersonalBar from './PersonalBar'

const Widget = ({info}) => (
  <Navbar inverse fixedTop fluid>
    <Navbar.Header>
      <Navbar.Brand>
        <Link to="/">{info.subTitle}</Link>
      </Navbar.Brand>
      <Navbar.Toggle />
    </Navbar.Header>
    <Navbar.Collapse>
      <Nav>
        {info.bottomLinks.map(function (l, i) {
          return (
            <LinkContainer key={i} to={l.href}>
              <NavItem eventKey={i}>
              {l.label}
              </NavItem>
            </LinkContainer>
          )
        })}
        <SwitchLang />
      </Nav>
      <Nav pullRight>
        <PersonalBar/>
      </Nav>
    </Navbar.Collapse>
  </Navbar>
)

Widget.propTypes = {
  info: PropTypes.object.isRequired
}

export default connect(
  state => ({info: state.siteInfo}),
  dispatch => ({})
)(Widget)
