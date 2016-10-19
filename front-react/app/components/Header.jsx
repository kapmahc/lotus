import React from 'react'
import {Navbar, Nav, NavItem} from 'react-bootstrap'

import SwitchLang from './SwitchLang'
import PersonalBar from './PersonalBar'

const Widget = () => (
  <Navbar inverse fixedTop fluid>
    <Navbar.Header>
      <Navbar.Brand>
        <a href="#">site.sub_title</a>
      </Navbar.Brand>
      <Navbar.Toggle />
    </Navbar.Header>
    <Navbar.Collapse>
      <Nav>
        <NavItem eventKey={1} href="#">Link</NavItem>
        <NavItem eventKey={2} href="#">Link</NavItem>
        <SwitchLang />
      </Nav>
      <Nav pullRight>
        <NavItem eventKey={1} href="#">Link Right</NavItem>
        <NavItem eventKey={2} href="#">Link Right</NavItem>
        <PersonalBar/>
      </Nav>
    </Navbar.Collapse>
  </Navbar>
)

export default Widget
