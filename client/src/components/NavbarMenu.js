import React from 'react'
import { Link } from 'react-router-dom'
import { Navbar, Form, Button } from 'react-bootstrap'

const NavbarMenu = () => {
	return (
		<div>
			<Navbar bg="dark" variant="dark">
				<Navbar.Brand as={Link} to="/">
					Hospital Playlist
				</Navbar.Brand>
				<Nav className="mr-auto">
					<Nav.Link as={Link} to="/">
						Home
					</Nav.Link>
				</Nav>
				<Form inline>
					<Button variant="outline-info">Login</Button>
					<Button variant="outline-info">Register</Button>
				</Form>
			</Navbar>
		</div>
	)
}

export default NavbarMenu
