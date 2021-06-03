import React from 'react'
import { Link, useHistory } from 'react-router-dom'
import { Form, Container, Alert } from 'react-bootstrap'

import { useSelector, useDispatch } from 'react-redux'

import userLoginAction from '../redux/actions/userLogin'

const Login = () => {
	const userLoginData = useSelector((state) => state.userLogin)
	const dispatch = useDispatch()
	const history = useHistory()

	const loginSubmitHandler = (e) => {
		e.preventDefault()

		dispatch(userLoginAction.login(userLoginData.email, userLoginData.password, history))
	}

	return (
		<>
			<Link to="/">Home</Link>
			<Container>
				<pre>{JSON.stringify(userLoginData)}</pre>
				{/* Error Message */}
				{userLoginData.errorMessage && (
					<Alert variant="danger">
						<ul>
							{userLoginData.errorMessage.map((error, index) => {
								return <li key={index}>{error}</li>
							})}
						</ul>
					</Alert>
				)}
				<Form onSubmit={loginSubmitHandler}>
					<Form.Group controlId="exampleForm.ControlInput1">
						<Form.Label>Email</Form.Label>
						<Form.Control
							type="email"
							placeholder="Email"
							required
							value={userLoginData.email}
							onChange={(e) => {
								dispatch(userLoginAction.setEmail(e.target.value))
							}}
						/>
					</Form.Group>
					<Form.Group controlId="exampleForm.ControlInput1">
						<Form.Label>Password</Form.Label>
						<Form.Control
							type="password"
							placeholder="Password"
							required
							value={userLoginData.password}
							onChange={(e) => {
								dispatch(userLoginAction.setPassword(e.target.value))
							}}
						/>
					</Form.Group>

					<Form.Group controlId="exampleForm.ControlTextarea1" className="mt-3">
						<Form.Control
							type="submit"
							value={userLoginData.isLoading ? 'Loading...' : 'Login'}
							disabled={userLoginData.isLoading ? true : false}
						/>
					</Form.Group>
				</Form>
			</Container>
		</>
	)
}

export default Login
