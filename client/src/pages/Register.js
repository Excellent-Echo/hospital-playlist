import React, { useEffect, useState } from 'react'
import { Form, Alert, Col } from 'react-bootstrap'
import DatePicker from 'react-datepicker'

import { useSelector, useDispatch } from 'react-redux'
import { Link } from 'react-router-dom'

import userRegisterAction from '../redux/actions/userRegister'

const RegisterPage = () => {
	const [startDate, setStartDate] = useState(new Date())
	console.log(startDate)
	const userRegisterData = useSelector((state) => state.userRegister)
	const dispatch = useDispatch()

	//  page pertama kali di load
	useEffect(() => {
		dispatch(userRegisterAction.resetForm())
	}, [])

	const handleRegisterSubmit = (e) => {
		e.preventDefault()
		dispatch(
			userRegisterAction.register(
				userRegisterData.email,
				userRegisterData.password,
				userRegisterData.fullName,
				userRegisterData.birthDate,
				userRegisterData.address,
				userRegisterData.gender
			)
		)
	}

	return (
		<>
			<h1>Register</h1>
			<div class="container col-md-4">
				<pre>{JSON.stringify(userRegisterData)}</pre>
				{/* Error Message */}
				{userRegisterData.errorMessage && (
					<Alert variant="danger">
						<ul>
							{userRegisterData.errorMessage.map((error, index) => {
								return <li key={index}>{error}</li>
							})}
						</ul>
					</Alert>
				)}
				{/* Success Message */}
				{userRegisterData.successMessage && <Alert variant="primary">{userRegisterData.successMessage}</Alert>}
				<Form onSubmit={handleRegisterSubmit}>
					<Form.Row>
						<Form.Group as={Col} controlId="formGridEmail">
							<Form.Label>Email</Form.Label>
							<Form.Control
								type="email"
								placeholder="Enter email"
								required
								// value={userRegisterData.email}
								onChange={(e) => dispatch(userRegisterAction.setEmail(e.target.value))}
							/>
						</Form.Group>

						<Form.Group as={Col} controlId="formGridPassword">
							<Form.Label>Password</Form.Label>
							<Form.Control
								type="password"
								placeholder="Password"
								required
								//   value={userRegisterData.password}
								onChange={(e) => dispatch(userRegisterAction.setPassword(e.target.value))}
							/>
						</Form.Group>
					</Form.Row>

					<Form.Group controlId="formFullName">
						<Form.Label>Full Name</Form.Label>
						<Form.Control
							type="text"
							placeholder="Your Name"
							required
							value={userRegisterData.fullName}
							onChange={(e) => dispatch(userRegisterAction.setFullName(e.target.value))}
						/>
					</Form.Group>

					<Form.Row>
						<Form.Group as={Col} controlId="formBirthDate">
							<Form.Label>Birth Date</Form.Label>
							<DatePicker
								selected={startDate}
								onChange={(date) => dispatch(userRegisterAction.setBirthDate(date))}
							/>
						</Form.Group>

						<Form.Group as={Col} controlId="formGender">
							<Form.Label>Gender</Form.Label>
							<Form.Control
								as="select"
								value={userRegisterData.gender}
								onChange={(e) => dispatch(userRegisterAction.setGender(e.target.value))}
							>
								<option value="Male">Male</option>
								<option value="Female">Female</option>
							</Form.Control>
						</Form.Group>
					</Form.Row>

					<Form.Group controlId="formGridAddress1">
						<Form.Label>Address</Form.Label>
						<Form.Control
							placeholder="1234 Main St"
							required
							value={userRegisterData.address}
							onChange={(e) => dispatch(userRegisterAction.setAddress(e.target.value))}
						/>
					</Form.Group>

					<Form.Group controlId="formSubmit">
						<Form.Control
							type="submit"
							value={userRegisterData.isLoading ? 'Loading...' : 'Register'}
							disabled={userRegisterData.isLoading ? true : false}
						/>
					</Form.Group>
				</Form>

				<Link to="/">Home</Link>
			</div>
		</>
	)
}

export default RegisterPage
