import React from 'react'
import {NavDropdown, MenuItem} from 'react-bootstrap'

const Widget = () => (
  <NavDropdown eventKey={3} title="Swtich lang" id="switch-lang-bar">
    <MenuItem eventKey={3.1}>English</MenuItem>
    <MenuItem eventKey={3.2}>简体中文</MenuItem>
    <MenuItem eventKey={3.3}>正體中文</MenuItem>
  </NavDropdown>
)

export default Widget
