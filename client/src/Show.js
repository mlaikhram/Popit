import { Button, Card, CardBody, CardTitle, Col, Collapse, Input, InputGroup, Nav, Navbar, NavbarBrand, NavbarToggler, NavItem, NavLink, Row, Spinner, UncontrolledCollapse } from "reactstrap"
import { PageCard } from "./components/PageCard";
import { PageScrollMenu } from "./components/PageScrollMenu";
import { Link, useParams } from "react-router-dom";

import logo from './logo.svg';
import './App.css';


function Show() { 
  const params = useParams();

  const show =     {
    "id": "625798abc15bd223b8a9eeae",
    "name": "Attack on Titan",
    "aliases": [
        "Shingeki no Kyojin"
    ],
    "synopsis": "Humanity struggles to fight back against the titans.",
    "images": {
        "1": "https://cdn.myanimelist.net/images/anime/5/44560l.jpg"
    }
}

  return (
    <div className="App">
      <Navbar color="dark" dark expand="md" fixed="top">
          <NavbarBrand href="/">Popit</NavbarBrand>
          <NavbarToggler onClick={function noRefCheck(){}} />
          <Collapse navbar>
              <Nav className="me-auto" navbar>
                  <NavItem>
                      <NavLink href={"/show/" + params.show}>{show.name}</NavLink>
                  </NavItem>
              </Nav>
          </Collapse>
      </Navbar>
    </div>
  );
}

export default Show;
