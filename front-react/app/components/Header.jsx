import React, { PropTypes } from 'react'
import { connect } from 'react-redux'
import {Navbar, Nav, NavItem} from 'react-bootstrap'

import SwitchLang from './SwitchLang'
import PersonalBar from './PersonalBar'

const Widget = ({info}) => (
  <Navbar inverse fixedTop fluid>
    <Navbar.Header>
      <Navbar.Brand>
        <a href="/">{info.subTitle}</a>
      </Navbar.Brand>
      <Navbar.Toggle />
    </Navbar.Header>
    <Navbar.Collapse>
      <Nav>
        {info.bottomLinks.map(function (l, i) {
          return (<NavItem eventKey={i} key={i} href={l.href}>
            {l.label}
          </NavItem>)
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
