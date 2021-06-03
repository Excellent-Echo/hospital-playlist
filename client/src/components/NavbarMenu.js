import React, { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { Nav, Navbar, Form, Button } from 'react-bootstrap'

import { useSelector } from 'react-redux'

const NavbarMenu = () => {
	const userLoginData = useSelector((state) => state.userLogin)
	// var x = localStorage.getItem('isAuth')
	// console.log(typeof x)

	const [isAuth, setIsAuth] = useState(false)

	useEffect(() => {
		localStorage.getItem('isAuth') === 'true' ? setIsAuth(true) : setIsAuth(false)
	})

	if (isAuth === true) {
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
						<Button variant="outline-info" as={Link} to="/">
							Logout
						</Button>
					</Form>
				</Navbar>
			</div>
		)
	}
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
					<Button variant="outline-info" as={Link} to="/login">
						Login
					</Button>
					<Button variant="outline-info" as={Link} to="/register">
						Register
					</Button>
				</Form>
			</Navbar>
		</div>
	)
}

export default NavbarMenu
