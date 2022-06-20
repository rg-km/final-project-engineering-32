import React, {useState} from 'react';
import { Navbar,Nav,NavDropdown,Container,Form,FormControl,Button } from 'react-bootstrap'
import Image from 'react-bootstrap/Image'
import { Link } from "react-router-dom"

export default function NavbarComponent() {
    
    return(
        <Navbar bg="light" expand="lg" position= "absolute">
        <Container>
            <Navbar.Brand href="#home">Budaya Baca</Navbar.Brand>
            <Navbar.Toggle aria-controls="basic-navbar-nav" />
            <Navbar.Collapse id="basic-navbar-nav">
            <Nav className="me-auto">
                <Nav.Link as={Link} to="/">Home</Nav.Link>
                <Nav.Link as={Link} to="/pinjamBuku">Pinjam Buku</Nav.Link>
                <NavDropdown title="Kategori" id="basic-nav-dropdown">
                <NavDropdown.Item href="#action/3.1">SD</NavDropdown.Item>
                <NavDropdown.Item href="#action/3.2">SMP</NavDropdown.Item>
                <NavDropdown.Item href="#action/3.3">SMA</NavDropdown.Item>
                <NavDropdown.Divider />
                <NavDropdown.Item href="#action/3.4">Umum</NavDropdown.Item>
                </NavDropdown>
            </Nav>
            <Form>
                <FormControl type="text" placeholder="Search" className="mr-sm-2" />
            </Form>
            <Button variant="outline-success">Search</Button>
            <Image src="../public/logo512.png" roundedCircle />
            </Navbar.Collapse>
        </Container>
        </Navbar>
    )
}
